//
//  main.swift
//  day2
//
//  Created by kamillo on 03/12/2020.
//

import Foundation

func isPassPhraseValid(passPhrase: String) -> Bool {
    var set: Set<String> = []
    
    return passPhrase.split(separator: " ").reduce(true, { (result, string) in
        if set.contains(String(string)) {
            return result && false
        } else {
            set.insert(String(string))
            return result && true
        }
    })
}


func isPassPhraseValid2(passPhrase: String) -> Bool {
    var set: Set<String> = []
    
    return passPhrase.split(separator: " ").reduce(true, { (result, string) in
        let sortedString = String(string.sorted())
        if set.contains(sortedString) {
            return result && false
        } else {
            set.insert(sortedString)
            return result && true
        }
    })
}

if let lines = getLines(fromFile: "input.txt") {
    var validCount = lines.reduce(0, { (result, phrase) in
        return isPassPhraseValid(passPhrase: String(phrase)) ? result + 1 : result
    })
    print("Part 1: ", validCount)
    
    validCount = lines.reduce(0, { (result, phrase) in
        return isPassPhraseValid2(passPhrase: String(phrase)) ? result + 1 : result
    })
    print("Part 2: ", validCount)
}
