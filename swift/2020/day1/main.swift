//
//  day1.swift
//  aoc-swift
//
//  Created by kamillo on 03/12/2020.
//

import Foundation

if let lines = getLines(fromFile: "day1.txt") {
    let numbers: [Int] = lines.compactMap{ Int($0) }
    let part1 = numbers.compactMap { num -> [Int]? in
        if let num2 = numbers.first(where: { num + $0 == 2020 }) {
            return [num, num2]
        }
        return nil
    }
    .first!
    .reduce(1){ $0 * $1 }
    
    let part2 = numbers.compactMap { num -> [[Int]]? in
        let arr = numbers.compactMap { num2 -> [Int]? in
            if let num3 = numbers.first(where: { num + num2 + $0 == 2020 }) {
                return [num, num2, num3]
            }
            return nil
        }
        if arr.isEmpty {
            return nil
        }
        
        return arr
    }
    .first!
    .first!
    .reduce(1){ $0 * $1 }
    
    print("Part 1: ", part1)
    print("Part 2: ", part2)
}
