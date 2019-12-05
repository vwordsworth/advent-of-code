from collections import Counter


def main():
    range_start = 172930
    range_end = 683083
    passwords = []

    for password in range(range_start, range_end):
        digits = [int(d) for d in str(password)]

        last_digit = digits[0]
        has_double_digits = False
        is_increasing = True

        for num in digits[1:]:
            if last_digit == num:
                has_double_digits = True        
            if last_digit > num:
                is_increasing = False
                break
            last_digit = num

        if has_double_digits and is_increasing:
            passwords.append(digits)

    num_meet_criteria = len(passwords)
    print("Part 1:\n\tPassword count = {0}".format(str(num_meet_criteria)))

    for password_digits in passwords:
        counts = Counter(password_digits)
        if 2 not in counts.values():
            num_meet_criteria -= 1
    print("Part 2:\n\tUpdated password count = {0}".format(str(num_meet_criteria)))


if __name__ == "__main__":
    main()
