# Assignment 10.1

# Python Practice: Working with Dictionaries
# For each problem, write a function that meets the requirements specified. Use the assert statements to test your code by uncommenting them.

# 1. 'add_key_value': Add a key-value pair to a dictionary
# Accepts a dictionary, a key, and a value
# Returns the updated dictionary

# assert add_key_value({"a": 1}, "b", 2) == {"a": 1, "b": 2}
# assert add_key_value({}, "key", "value") == {"key": "value"}

# 2. 'remove_key': Remove a key from a dictionary
# Accepts a dictionary and a key
# Returns the updated dictionary (or the original dictionary if the key is not found)

# assert remove_key({"a": 1, "b": 2}, "b") == {"a": 1}
# assert remove_key({"a": 1}, "c") == {"a": 1}

# 3. 'key_exists': Check if a key exists in a dictionary
# Accepts a dictionary and a key
# Returns a boolean

# assert key_exists({"a": 1, "b": 2}, "a") == True
# assert key_exists({"a": 1, "b": 2}, "c") == False

# 4. 'value_exists': Check if a value exists in a dictionary
# Accepts a dictionary and a value
# Returns a boolean

# assert value_exists({"a": 1, "b": 2}, 2) == True
# assert value_exists({"a": 1, "b": 2}, 3) == False

# 5. 'get_keys': Return all keys in a dictionary as a list
# Accepts a dictionary
# Returns a list

# assert get_keys({"a": 1, "b": 2}) == ["a", "b"]
# assert get_keys({}) == []

# 6. 'get_values': Return all values in a dictionary as a list
# Accepts a dictionary
# Returns a list

# assert get_values({"a": 1, "b": 2}) == [1, 2]
# assert get_values({}) == []

# 7. 'merge_dicts': Merge two dictionaries, combining their key-value pairs
# Accepts two dictionaries
# Returns a dictionary

# assert merge_dicts({"a": 1}, {"b": 2}) == {"a": 1, "b": 2}
# assert merge_dicts({}, {}) == {}

# 8. 'invert_dict': Invert a dictionary so that keys become values and values become keys
# Accepts a dictionary
# Returns a dictionary

# assert invert_dict({"a": 1, "b": 2}) == {1: "a", 2: "b"}
# assert invert_dict({}) == {}

# 9. 'sum_dict_values': Return the sum of all numeric values in a dictionary
# Accepts a dictionary with numeric values
# Returns a number

# assert sum_dict_values({"a": 1, "b": 2, "c": 3}) == 6
# assert sum_dict_values({}) == 0

# 10. 'dict_from_keys': Create a dictionary with specified keys and a default value
# Accepts a list of keys and a default value
# Returns a dictionary

# assert dict_from_keys(["a", "b", "c"], 0) == {"a": 0, "b": 0, "c": 0}
# assert dict_from_keys([], "default") == {}
