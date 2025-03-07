const std = @import("std");
const print = std.debug.print;

pub fn main() !void {
    const file = try std.fs.cwd().openFile("data/2015/day4.txt", .{});
    defer file.close();

    var gpa = std.heap.page_allocator;

    const content = try file.readToEndAlloc(gpa, 9);
    defer gpa.free(content);

    var suffix: i32 = 0;

    var hasher = std.crypto.hash.Md5.init(.{});
    var digest: [16]u8 = undefined;
    var buf: [16]u8 = undefined;
    var done: bool = false;

    while (!done) {
        suffix += 1;
        const suffix_as_string = try std.fmt.bufPrint(&buf, "{}", .{suffix});

        var buffer: [24]u8 = undefined;

        @memcpy(buffer[0..content.len], content);
        @memcpy(buffer[content.len - 1 .. content.len + suffix_as_string.len - 1], suffix_as_string);

        const concatenated = buffer[0..(content.len + suffix_as_string.len - 1)];

        hasher = std.crypto.hash.Md5.init(.{});
        hasher.update(concatenated);
        hasher.final(&digest);
        const hash = std.fmt.bytesToHex(digest, .lower);

        for (hash[0..6]) |byte| {
            if (byte != '0') {
                done = false;
                break;
            }
            done = true;
        }
    }

    print("First number to get 5 leading zeros: {d}\n", .{suffix});
}
