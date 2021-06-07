from typing import Text
import errno
import os
import io

def create_dir(dir_path:str):
    """Creates a directory and its super paths.
    Succeeds even if the path already exists."""

    try:
        os.makedirs(dir_path)
    except OSError as e:
        # be happy if someone already created the path
        if e.errno != errno.EEXIST:
            raise

def create_dir_for_file(file_path):
    # type: (Text) -> None
    """Creates any missing parent directories of this files path."""

    try:
        os.makedirs(os.path.dirname(file_path))
    except OSError as e:
        # be happy if someone already created the path
        if e.errno != errno.EEXIST:
            raise

def write_to_file(filename, text, auto_create_dir=False):
    # type: (Text, Text, bool) -> None
    """Write a text to a file."""
    from os.path import expanduser
    filename=expanduser(filename)
    if auto_create_dir:
        create_dir_for_file(filename)

    with io.open(filename, 'w', encoding="utf-8") as f:
        f.write(str(text))
