const std = @import("std");
const print = std.debug.print;
const cal = 52;

pub fn main() !void {
    const file = try std.fs.cwd().openFile("data/2022/day3.txt", .{});
    defer file.close();

    const reader = file.reader();
    var buf: [1024]u8 = undefined;

    var dupe_cnt: u32 = 0;
    var curr_chars: [cal]u8 = undefined;
    var line: u8 = 0;
    var second: u32 = 0;

    while (reader.readUntilDelimiterOrEof(&buf, '\n')) |maybe_line| {
        if (maybe_line == null) break;

        dupe_cnt += first(maybe_line.?);

        if (line == 0) {
            curr_chars = getChars(maybe_line.?);
        } else {
            for (curr_chars, 0..) |char, idx| {
                if (contains(u8, maybe_line.?, char)) {
                    continue;
                }
                curr_chars[idx] = 0;
            }
        }

        if (line == 2) {
            line = 0;
            const eva = evaulateDupes(curr_chars);
            second += eva;
            continue;
        }
        line += 1;
    } else |err| {
        print("Error while reading the file: {}\n", .{err});
    }

    print("First: {}\n", .{dupe_cnt});
    print("Second: {}\n", .{second});
}

fn first(line: []u8) u32 {
    const l = line.len >> 1;
    const dupes: [cal]u8 = getDupes(line[0..l], line[l..]);
    return evaulateDupes(dupes);
}

fn evaulateDupes(dupes: [cal]u8) u32 {
    var res: u32 = 0;
    for (dupes) |dupe| {
        if (dupe == 0) {
            continue;
        }
        if (dupe > 90) {
            res += 1 + dupe - 'a';
        } else {
            res += 27 + dupe - 'A';
        }
    }
    return res;
}

fn getDupes(front: []u8, back: []u8) [cal]u8 {
    var result: [cal]u8 = [1]u8{0} ** cal;
    var idx: usize = 0;
    for (front) |char| {
        if (contains(u8, &result, char)) {
            continue;
        }
        if (contains(u8, back, char)) {
            result[idx] = char;
            idx += 1;
            continue;
        }
    }
    return result;
}

fn getChars(line: []u8) [cal]u8 {
    var res: [cal]u8 = [1]u8{0} ** cal;
    var idx: usize = 0;
    for (line) |char| {
        if (contains(u8, &res, char)) {
            continue;
        }
        res[idx] = char;
        idx += 1;
    }
    return res;
}

fn contains(comptime T: type, arr: []T, val: T) bool {
    for (arr) |ch| {
        if (ch == val) {
            return true;
        }
    }
    return false;
}
