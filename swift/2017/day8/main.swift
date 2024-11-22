//
//  main.swift
//  day8
//
//  Created by kamillo on 03/12/2020.
//

import Foundation

if let lines = getLines(fromFile: "input.txt") {
    var registers: [String:Int] = [:]
    var max = 0
    
    lines.forEach { line in
        let split = line.split(separator: " ")
        
        let condReg = String(split[4])
        let cond = String(split[5])
        let condVal = Int(split[6])!
        
        var condResult = false
        switch cond {
        case "<" : condResult = registers[condReg, default: 0] < condVal; break
        case ">" : condResult = registers[condReg, default: 0] > condVal; break
        case "==" : condResult = registers[condReg, default: 0] == condVal; break
        case "<=" : condResult = registers[condReg, default: 0] <= condVal; break
        case ">=" : condResult = registers[condReg, default: 0] >= condVal; break
        case "!=" : condResult = registers[condReg, default: 0] != condVal; break
        default: break
        }
        
        if condResult {
            let reg = String(split[0])
            
            if let val = Int(split[2]) {
                if String(split[1]) == "inc" {
                    registers[reg, default: 0] += val
                } else {
                    registers[reg, default: 0] -= val
                }
            }
        }
        
        if let m = registers.max(by: { a, b in a.value < b.value }), m.value > max {
            max = m.value
        }
    }
    
    print("Part 1:", registers.max { a, b in a.value < b.value }!.value)
    print("Part 2:", max)
}
