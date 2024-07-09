const std = @import("std");
const print = std.debug.print;


pub fn main() void {
    const file = std.fs.cwd().openFile("../../../data/2022/day2.txt", .{}) catch |err| {
        print("Could not open file: {}\n", .{err});
        std.process.exit(1);
    };

    const list = parseFile(file);

    for (list.items) |item| {
        print("{s}\n", .{item});
    }
}


fn parseFile(file: std.fs.File) std.ArrayList([]u8) {
    const reader = file.reader();
    var buff: [1024]u8 = undefined;

    const allocator = std.heap.page_allocator;
    var result = std.ArrayList([]u8).init(allocator);

    while (reader.readUntilDelimiterOrEof(&buff, '\n')) |maybe_line| {
        if (maybe_line == null) {
            break;
        }
        result.append(maybe_line.?) catch |err| {
            print("Could not append line: {}\n", .{err});
        };
    } else |err| {
        print("Error while reading file: {}\n", .{err});
    }

    return result;
}
