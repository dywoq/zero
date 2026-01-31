// Copyright 2026 dywoq
//
// This code is released under the Apache License 2.0:
// https://www.apache.org/licenses/LICENSE-2.0
//
// Repository:
// https://github.com/dywoq/zero

/// Represents the snapshot time, received from the bootloader.
#[repr(C)]
pub struct SnapshotTime {
    pub year: u16,
    pub month: u8,
    pub day: u8,
    pub hour: u8,
    pub minute: u8,
    pub second: u8,
    pub nanoseconds: u32,
}

impl SnapshotTime {
    /// Returns a instance of [SnapshotTime] with time, received from [uefi::runtime::Time].
    pub fn from_uefi(time: uefi::runtime::Time) -> SnapshotTime {
        SnapshotTime {
            year: time.year(),
            month: time.month(),
            day: time.day(),
            hour: time.hour(),
            minute: time.minute(),
            second: time.second(),
            nanoseconds: time.nanosecond(),
        }
    }
}
