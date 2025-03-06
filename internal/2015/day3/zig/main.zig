const std = @import("std");
const fs = std.fs;
const print = std.debug.print;

pub fn main() !void {
    const file = try fs.cwd().openFile("data/2015/day3.txt", .{});
    defer file.close();

    const reader = file.reader();
    var buf: [8193]u8 = undefined;

    while (reader.readUntilDelimiterOrEof(&buf, '\n')) |maybe_line| {
        if (maybe_line == null) break;
        print("{s}\n", .{maybe_line.?});
    } else |err| {
        print("Could not open file {}\n", .{err});
    }
}
