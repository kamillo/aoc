//
//  Permutations.swift
//  aoc-swift
//
//  Created by kamillo on 03/12/2020.
//

import Foundation

extension Array {
    func decompose() -> (Iterator.Element, [Iterator.Element])? {
        guard let x = first else { return nil }
        return (x, Array(self[1..<count]))
    }
}

func between<T>(x: T, _ ys: [T]) -> [[T]] {
    guard let (head, tail) = ys.decompose() else { return [[x]] }
    return [[x] + ys] + between(x: x, tail).map { [head] + $0 }
}

func permutations<T>(xs: [T]) -> [[T]] {
    guard let (head, tail) = xs.decompose() else { return [[]] }
    return permutations(xs: tail).flatMap { between(x: head, $0) }
}
