
proc getLines*(file: string): seq[string] =
    var 
        input: File
        line: string

    if open(input, file):
        while input.readLine(line):
            result.add(line)

