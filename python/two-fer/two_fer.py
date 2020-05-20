def two_fer(name=None):
    if name is None:
        return "One for you, one for me."

    return "One for %s, one for me." % name
