const std = @import("std");
const print = std.debug.print;


pub fn main() void {
    const file = std.fs.cwd().openFile("../../../data/2022/day2.txt", .{}) catch |err| {
        print("Could not open file: {}\n", .{err});
        std.process.exit(1);
    };

    const list = parseFile(file);

    print("First: {d}\n", .{first(list)});
    print("Second: {d}\n", .{second(list)});
}

fn first(list: std.ArrayList([2]u8)) u16 {
    var score: u16 = 0;
    for (list.items) |item| {
        switch (item[0]) {
            'A' => switch (item[1]) {
                'X' => score += 4,
                'Y' => score += 8,
                'Z' => score += 3,
                else => unreachable,
            },
            'B' => switch (item[1]) {
                'X' => score += 1,
                'Y' => score += 5,
                'Z' => score += 9,
                else => unreachable,
            },
            'C' => switch (item[1]) {
                'X' => score += 7,
                'Y' => score += 2,
                'Z' => score += 6,
                else => unreachable,
            },
            else => unreachable,
        }
    }
    return score;
}

fn second(list: std.ArrayList([2]u8)) u16 {
    var score: u16 = 0;
    for (list.items) |item| {
        switch (item[0]) {
            'A' => switch (item[1]) {
                'X' => score += 3,
                'Y' => score += 4,
                'Z' => score += 8,
                else => unreachable,
            },
            'B' => switch (item[1]) {
                'X' => score += 1,
                'Y' => score += 5,
                'Z' => score += 9,
                else => unreachable,
            },
            'C' => switch (item[1]) {
                'X' => score += 2,
                'Y' => score += 6,
                'Z' => score += 7,
                else => unreachable,
            },
            else => unreachable,
        }
    }
    return score;
}

fn parseFile(file: std.fs.File) std.ArrayList([2]u8) {
    const reader = file.reader();
    var buff: [1024]u8 = undefined;

    const allocator = std.heap.page_allocator;
    var result = std.ArrayList([2]u8).init(allocator);

    while (reader.readUntilDelimiterOrEof(&buff, '\n')) |maybe_line| {
        if (maybe_line == null) {
            break;
        }
        result.append(maybe_line.?[0..1].* ++ maybe_line.?[2..3].*) catch |err| {
            print("Could not append line: {}\n", .{err});
        };
    } else |err| {
        print("Error while reading file: {}\n", .{err});
    }

    return result;
}
