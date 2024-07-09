const std = @import("std");
const print = std.debug.print;

pub fn main() void {
    const file_path = "../../../data/2022/day1.txt";

    const file: std.fs.File = std.fs.cwd().openFile(file_path, .{}) catch |err| {
        print("An error occurred while trying to open file at '{s}': {}\n", .{file_path, err});
        std.process.exit(1);
    };
    defer file.close();

    const values: std.ArrayList(i32) = parseFile(file);
    defer values.deinit();

    print("First: {}\n", .{ first(values) });
    print("Second: {}\n", .{ second(values) });
}

pub fn parseFile(file: std.fs.File) std.ArrayList(i32) {
    const reader = file.reader();
    var buff: [1024]u8 = undefined;

    const allocator = std.heap.page_allocator;
    var res = std.ArrayList(i32).init(allocator); 

    var curr: i32 = 0;

    while (reader.readUntilDelimiterOrEof(&buff, '\n')) |maybe_line| {
        if (maybe_line == null) {
            res.append(curr) catch |err| {
                print("Could not append to array list: {}\n", .{err});
            };
            curr = 0;
            break;
        }
        if (maybe_line.?.len == 0) {
            res.append(curr) catch |err| {
                print("Could not append to array list: {}\n", .{err});
            };
            curr = 0;
        } else {
            curr += std.fmt.parseInt(i32, maybe_line.?, 10) catch |err| {
                print("Error parsing {s}: {}\n", .{maybe_line.?, err});
                std.process.exit(1);
            };
        }
    } else |err| {
        print("an error occurred while reading line: {}\n", .{err});
    }
    return res;
}

pub fn first(values: std.ArrayList(i32)) i32 {
    var max: i32 = 0;
    for (values.items) |value| {
        if (value > max) {
            max = value;
        }
    }
    return max;
}

pub fn second(values: std.ArrayList(i32)) i32 {
    var max = [3]i32{0, 0, 0};
    for (values.items) |value| {
        if (value > max[0]) {
            max[2] = max[1];
            max[1] = max[0];
            max[0] = value;
        } else if (value > max[1]) {
            max[2] = max[1];
            max[1] = value;
        } else if (value > max[2]) {
            max[2] = value;
        }
    }
    return max[0] + max[1] + max[2];
}
