python_library(
  name = "conway",
  srcs = ["conway.py"],
)

python_test(
  name = "conway_test",
  srcs = ["conway_test.py"],
  deps = [":conway"],
)

python_binary(
  name = "main",
  main = "main.py",
  deps = [":conway"],
)
