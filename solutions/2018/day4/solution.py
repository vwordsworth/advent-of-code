from guard import Guard
from log import Log
from log_type import LogTypes
from operator import attrgetter, itemgetter


def main():
    logs = read_input()
    all_guards = get_all_guards(logs)
    guard = find_sleepiest_guard(all_guards)
    minute = find_sleepiest_minute(guard)
    result = int(guard.id) * minute
    print("Part 1:\n\tGuard ID ({0}) * Minute Chosen ({1}): {2}".format(guard.id, str(minute), str(result)))
    guard = find_most_frequently_sleepy_guard(all_guards)
    minute = guard.most_frequent_minute
    result = int(guard.id) * minute
    print("Part 2:\n\tGuard ID ({0}) * Minute Chosen ({1}): {2}".format(guard.id, str(minute), str(result)))


def read_input():
    return [line.rstrip('\n') for line in open('data/sorted.txt')]


def get_all_guards(logs):
    guards = {}
    current_sleep_start = None
    current_guard = None

    for entry in logs:
        log = Log(entry)
        if log.log_type == LogTypes.NEW_GUARD:
            guard_id = log.get_guard_number_from_new_guard_message()
            if guard_id in guards:
                current_guard = guards[guard_id]
            else:
                new_guard = Guard(guard_id)
                guards[guard_id] = new_guard
                current_guard = new_guard
        elif log.log_type == LogTypes.FALL_ASLEEP:
            current_sleep_start = log.minute
        elif log.log_type == LogTypes.WAKE_UP:
            current_guard.add_sleep(int(current_sleep_start), int(log.minute))
    
    return list(guards.values())


def find_sleepiest_guard(all_guards):
    all_guards.sort(key=attrgetter("total_sleep_time"), reverse=True)
    return all_guards[0]


def find_sleepiest_minute(guard):
    return max(guard.minute_frequency.items(), key=itemgetter(1))[0]


def find_most_frequently_sleepy_guard(all_guards):
    all_guards.sort(key=attrgetter("most_frequent_count"), reverse=True)
    return all_guards[0]


if __name__ == "__main__":
    main()
