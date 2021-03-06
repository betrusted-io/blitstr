// Copyright (c) 2020 Sam Blenny
// SPDX-License-Identifier: Apache-2.0 OR MIT
//
#![forbid(unsafe_code)]

/// Point specifies a pixel coordinate
#[derive(Copy, Clone, Debug, PartialEq, PartialOrd, rkyv::Archive, rkyv::Unarchive)]
pub struct Pt {
    pub x: u32,
    pub y: u32,
}

impl Pt {
    /// Make a new point
    pub fn new(x: u32, y: u32) -> Pt {
        Pt { x, y }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_pt_equivalence() {
        let p1 = Pt { x: 1, y: 2 };
        let p2 = Pt::new(1, 2);
        assert_eq!(p1, p2);
    }

    #[test]
    fn test_pt_ordering() {
        let p1 = Pt { x: 1, y: 2 };
        let p2 = Pt::new(1, 3);
        let p3 = Pt::new(0, 0);
        assert!(p1 < p2);
        assert!(p1 > p3);
    }
}