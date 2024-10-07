// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @title TokenVesting
 * @dev 代币归属合约，处理多个受益人的归属计划。
 * 允许在适用的情况下撤销归属计划，并根据角色区分归属逻辑。
 */
contract TokenVesting is Ownable {
    // 被归属的ERC20代币
    IERC20 public token;

    // 角色定义
    enum Role {
        Investor,
        Employee
    }

    // 释放间隔定义
    enum ReleaseInterval {
        Minute,
        Hour,
        Day,
        Week,
        Month,
        Quarter,
        HalfYear,
        Year
    }

    // 表示归属计划的结构体
    struct VestingSchedule {
        bool initialized; // 归属计划是否已初始化
        address beneficiary; // 受益人地址
        uint256 cliff; // 悬崖期时长（秒）
        uint256 start; // 归属期开始时间
        uint256 duration; // 归属期时长
        uint256 slicePeriodSeconds; // 归属周期时长（秒）
        uint256 totalAmount; // 被归属的代币总量
        uint256 released; // 已释放的代币数量
        bool revocable; // 归属计划是否可撤销
        bool revoked; // 归属计划是否已被撤销
        Role role; // 受益人的角色
        uint256 cliffReleasePercentage; // 悬崖期到期后一次性释放的百分比
    }

    // 受益人地址到归属计划的映射
    mapping(address => VestingSchedule) private vestingSchedules;
    // 所有归属计划中的代币总量
    uint256 private vestingSchedulesTotalAmount;

    // 记录重要操作的事件
    event Released(address indexed beneficiary, uint256 amount);
    event Revoked(address indexed beneficiary, uint256 amount);
    event VestingScheduleCreated(
        address indexed beneficiary,
        uint256 totalAmount,
        Role role,
        uint256 start,
        uint256 duration,
        uint256 cliff,
        uint256 slicePeriodSeconds,
        uint256 cliffReleasePercentage
    );

    /**
     * @dev 构造函数，初始化合约并设置代币地址
     * @param token_ ERC20代币合约的地址
     */
    constructor(address token_) Ownable(msg.sender) {
        require(token_ != address(0x0), "Token address cannot be zero");
        token = IERC20(token_);
    }

    // 修饰符：检查受益人的归属计划是否存在
    modifier onlyIfVestingScheduleExists(address _beneficiary) {
        require(
            vestingSchedules[_beneficiary].initialized,
            "Vesting schedule does not exist"
        );
        _;
    }

    // 修饰符：检查归属计划是否未被撤销
    modifier onlyIfNotRevoked(address _beneficiary) {
        require(
            !vestingSchedules[_beneficiary].revoked,
            "Vesting schedule is revoked"
        );
        _;
    }

    /**
     * @dev 返回受益人的归属计划及当前可获取的额度
     * @param _beneficiary 受益人地址
     */
    function getVestingSchedule(
        address _beneficiary
    )
        public
        view
        onlyIfVestingScheduleExists(_beneficiary)
        returns (
            uint256 cliff,
            uint256 start,
            uint256 duration,
            uint256 slicePeriodSeconds,
            uint256 totalAmount,
            uint256 released,
            bool revocable,
            bool revoked,
            Role role,
            uint256 cliffReleasePercentage,
            uint256 releasableAmount
        )
    {
        VestingSchedule memory vesting = vestingSchedules[_beneficiary];
        uint256 releasable = _computeReleasableAmount(vesting);
        return (
            vesting.cliff,
            vesting.start,
            vesting.duration,
            vesting.slicePeriodSeconds,
            vesting.totalAmount,
            vesting.released,
            vesting.revocable,
            vesting.revoked,
            vesting.role,
            vesting.cliffReleasePercentage,
            releasable
        );
    }

    /**
     * @dev 创建受益人的归属计划
     * @param _beneficiary 受益人地址
     * @param _start 归属期开始时间
     * @param _cliff 悬崖期时长（秒）
     * @param _duration 归属期时长（秒）
     * @param _totalAmount 被归属的代币总量
     * @param _revocable 归属计划是否可撤销
     * @param _role 受益人的角色（Investor 或 Employee）
     * @param _releaseInterval 释放间隔（Minute, Hour, Day, Week, Month, Quarter, HalfYear, Year）
     * @param _cliffReleasePercentage 悬崖期到期后一次性释放的百分比（以百分之一为单位，如20%则为2000）
     */
    function createVestingSchedule(
        address _beneficiary,
        uint256 _start,
        uint256 _cliff,
        uint256 _duration,
        uint256 _totalAmount,
        bool _revocable,
        Role _role,
        ReleaseInterval _releaseInterval,
        uint256 _cliffReleasePercentage
    ) public onlyOwner {
        require(
            _beneficiary != address(0x0),
            "Beneficiary cannot be zero address"
        );
        require(_duration > 0, "Duration must be > 0");
        require(_totalAmount > 0, "Total amount must be > 0");
        require(_duration >= _cliff, "Duration must be >= cliff");
        require(
            _cliffReleasePercentage <= 10000,
            "Cliff release percentage must be <= 10000"
        );
        require(
            !vestingSchedules[_beneficiary].initialized,
            "Vesting schedule already exists for this beneficiary"
        );

        uint256 slicePeriodInSeconds = _getSlicePeriodSeconds(_releaseInterval);

        VestingSchedule memory schedule;
        schedule.initialized = true;
        schedule.beneficiary = _beneficiary;
        schedule.cliff = _start + _cliff;
        schedule.start = _start;
        schedule.duration = _duration;
        schedule.slicePeriodSeconds = slicePeriodInSeconds;
        schedule.totalAmount = _totalAmount;
        schedule.revocable = _revocable;
        schedule.role = _role;
        schedule.cliffReleasePercentage = _cliffReleasePercentage;
        vestingSchedules[_beneficiary] = schedule;
        vestingSchedulesTotalAmount =
            vestingSchedulesTotalAmount +
            _totalAmount;

        emit VestingScheduleCreated(
            _beneficiary,
            _totalAmount,
            _role,
            _start,
            _duration,
            _cliff,
            slicePeriodInSeconds,
            _cliffReleasePercentage
        );
    }

    /**
     * @dev 根据释放间隔计算归属周期的秒数
     * @param _releaseInterval 释放间隔
     * @return uint256 归属周期的秒数
     */
    function _getSlicePeriodSeconds(
        ReleaseInterval _releaseInterval
    ) internal pure returns (uint256) {
        if (_releaseInterval == ReleaseInterval.Minute) {
            return 60;
        } else if (_releaseInterval == ReleaseInterval.Hour) {
            return 3600;
        } else if (_releaseInterval == ReleaseInterval.Day) {
            return 86400;
        } else if (_releaseInterval == ReleaseInterval.Week) {
            return 604800;
        } else if (_releaseInterval == ReleaseInterval.Month) {
            return 2592000;
        } else if (_releaseInterval == ReleaseInterval.Quarter) {
            return 7776000; // 三个月
        } else if (_releaseInterval == ReleaseInterval.HalfYear) {
            return 15552000; // 六个月
        } else if (_releaseInterval == ReleaseInterval.Year) {
            return 31536000; // 一年
        } else {
            revert("Invalid release interval");
        }
    }

    /**
     * @dev 向受益人释放已归属的代币
     * @param _beneficiary 受益人地址
     */
    function release(
        address _beneficiary
    )
        public
        onlyIfVestingScheduleExists(_beneficiary)
        onlyIfNotRevoked(_beneficiary)
    {
        VestingSchedule storage vesting = vestingSchedules[_beneficiary];
        uint256 vestedAmount = _computeReleasableAmount(vesting);
        require(vestedAmount > 0, "No releasable amount available");

        vesting.released = vesting.released + vestedAmount;
        token.transfer(vesting.beneficiary, vestedAmount);
        emit Released(_beneficiary, vestedAmount);
    }

    /**
     * @dev 撤销可撤销的归属计划
     * @param _beneficiary 受益人地址
     */
    function revoke(
        address _beneficiary
    )
        public
        onlyOwner
        onlyIfVestingScheduleExists(_beneficiary)
        onlyIfNotRevoked(_beneficiary)
    {
        VestingSchedule storage vesting = vestingSchedules[_beneficiary];
        require(vesting.revocable, "Vesting schedule is not revocable");

        vesting.revoked = true;
        uint256 unreleased = vesting.totalAmount - vesting.released;
        vestingSchedulesTotalAmount = vestingSchedulesTotalAmount - unreleased;
        token.transfer(owner(), unreleased);
        emit Revoked(_beneficiary, unreleased);
    }

    /**
     * @dev 计算归属计划的可释放代币数量
     * @param vesting 归属计划
     */
    function _computeReleasableAmount(
        VestingSchedule memory vesting
    ) internal view returns (uint256) {
        if (block.timestamp < vesting.cliff) {
            return 0;
        } else if (vesting.revoked) {
            return 0;
        } else if (block.timestamp >= vesting.start + vesting.duration) {
            return vesting.totalAmount - vesting.released;
        } else {
            uint256 cliffReleaseAmount = (vesting.totalAmount *
                vesting.cliffReleasePercentage) / 10000;
            uint256 timeFromCliff = block.timestamp - vesting.cliff;
            uint256 vestedSlices = timeFromCliff / vesting.slicePeriodSeconds;
            uint256 vestedSeconds = vestedSlices * vesting.slicePeriodSeconds;
            uint256 vestedAmountAfterCliff = ((vesting.totalAmount -
                cliffReleaseAmount) * vestedSeconds) /
                (vesting.duration - (vesting.cliff - vesting.start));
            return
                cliffReleaseAmount + vestedAmountAfterCliff - vesting.released;
        }
    }

    /**
     * @dev 提取调用者的可释放金额
     */
    function withdraw() public {
        address beneficiary = msg.sender;
        require(
            vestingSchedules[beneficiary].initialized,
            "Vesting schedule does not exist"
        );
        require(
            !vestingSchedules[beneficiary].revoked,
            "Vesting schedule is revoked"
        );

        VestingSchedule storage vesting = vestingSchedules[beneficiary];
        uint256 vestedAmount = _computeReleasableAmount(vesting);
        require(vestedAmount > 0, "No releasable amount available");

        vesting.released = vesting.released + vestedAmount;
        token.transfer(beneficiary, vestedAmount);
        emit Released(beneficiary, vestedAmount);
    }

    /**
     * @dev 返回所有归属计划中的代币总量
     */
    function getVestingSchedulesTotalAmount() public view returns (uint256) {
        return vestingSchedulesTotalAmount;
    }

    /**
     * @dev 返回代币合约地址
     */
    function getTokenAddress() public view returns (address) {
        return address(token);
    }
}
