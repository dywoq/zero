// Copyright 2026 dywoq
//
// This code is released under the Apache License 2.0:
// https://www.apache.org/licenses/LICENSE-2.0
//
// Repository:
// https://github.com/dywoq/zero

use uefi::runtime;
use zero_boot::time::SnapshotTime;

pub fn get_time() -> SnapshotTime {
    let time = runtime::get_time().expect("Failed to get time");
    let snapshot = SnapshotTime::from_uefi(time);
    snapshot
}
