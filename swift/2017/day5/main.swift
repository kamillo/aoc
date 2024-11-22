//
//  main.swift
//  day2
//
//  Created by kamillo on 03/12/2020.
//

import Foundation

if let lines = getLines(fromFile: "input.txt") {
    var instructionArray = lines.map { Int(String($0))! }

    var i = 0
    var count = 0
    while true {
        if (i >= instructionArray.count || i < 0) {
            break;
        }
        
        let jump = instructionArray[i]
        instructionArray[i] += 1
        i += jump
        count += 1
    }
    
    print("Part 1: ", count)
    
    instructionArray = lines.map { Int(String($0))! }
    i = 0
    count = 0
    while true {
        if (i >= instructionArray.count || i < 0) {
            break;
        }
        
        let jump = instructionArray[i]
        instructionArray[i] += jump >= 3 ? -1 : 1
        i += jump
        count += 1
    }
    
    print("Part 2: ", count)
}
