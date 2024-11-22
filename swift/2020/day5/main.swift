//
//  main.swift
//  day2
//
//  Created by kamillo on 03/12/2020.
//

import Foundation

if let lines = getLines(fromFile: "day5.txt") {
//    let bin = lines.map{ $0.replacingOccurrences(of: "R", with: "1")}
    let ids = lines.map{ line -> Int in
        let converted = line.map{ char -> String.Element in
            switch char {
            case "R", "B": return "1"
            case "L", "F": return "0"
            default: return char
            }
        }
        
        let row = Int(String(converted[..<7]), radix: 2)!
        let col = Int(String(converted[7...]), radix: 2)!
        
        return row * 8 + col
    }.sorted()
    print(ids.last!)
    print(ids.enumerated().filter{ $0.offset > 0 && ($0.element - ids[$0.offset - 1] != 1)}.first!.element - 1)
}

