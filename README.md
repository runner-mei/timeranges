# Time Ranges Subtraction

Subtract one list of time ranges from another. Formally: for two
lists of time ranges A and B, a time is in (A-B) if and only if it is part of A and not part of
B. A time range has a start time and an end time.

Note: This algorithm only works when both A and B do NOT have overlaps in themselves.

## Example

```
(9:00-10:00) “minus” (9:00-9:30) = (9:30-10:00)
(9:00-10:00) “minus” (9:00-10:00) = ()
(9:00-9:30) “minus” (9:30-15:00) = (9:00-9:30)
(9:00-9:30, 10:00-10:30) “minus” (9:15-10:15) = (9:00-9:15, 10:15-10:30)
(9:00-11:00, 13:00-15:00) “minus” (9:00-9:15, 10:00-10:15, 12:30-16:00)
= (9:15-10:00, 10:15-11:00)
```

## Run test

```
git clone https://github.com/fenggaocloud/timeranges.git
cd timeranges/timerangesub
go test
```
 
 You should see the message below:

```
PASS
ok      github.com/fenggaocloud/timeranges/timerangesub 0.006s
```

## Add more test cases

Add new test cases in the test file ./timerangesub/timerangesub_test.go

## TODOs

. To add validation and overlap handling in each time ranges.
