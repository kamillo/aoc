//
//  main.swift
//  2017-day1
//
//  Created by kamillo on 04/12/2021.
//

import Foundation

if let lines = getLines(fromFile: "input.txt") {
    print("Part 1:", calc(seq: lines[0]))
    print("Part 2:", calc2(seq: lines[0]))
}

func calc(seq: String) -> Int {
    let array = Array(seq)
    return array.enumerated().reduce(0) { (x, y) -> Int in
        let current = Int(String(y.element))!
        let next = Int(String(array[(y.offset + 1) % array.count]))!
                
        return current == next ? x + current : x
    }
}

func calc2(seq: String) -> Int {
    let array = Array(seq)
    return array.enumerated().reduce(0) { (x, y) -> Int in
        let current = Int(String(y.element))!
        let next = Int(String(array[(y.offset + array.count/2) % array.count]))!
                
        return current == next ? x + current : x
    }
}
