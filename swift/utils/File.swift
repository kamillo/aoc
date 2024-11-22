//
//  File.swift
//  aoc-swift
//
//  Created by kamillo on 03/12/2020.
//

import Foundation

func getLines(fromFile file: String) -> [String]? {
    do {
        let contents = try String(contentsOfFile: file, encoding: .utf8)
        return contents.split(separator: "\n").map{String($0)}
    }
    catch let error as NSError {
        print("Error opening file: \(error)")
    }
    
    return nil
}
