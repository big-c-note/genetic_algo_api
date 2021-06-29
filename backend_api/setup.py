import setuptools


def get_package_description() -> str:
    """Returns a description of this package from the markdown files."""
    with open("README.md", "r") as stream:
        readme: str = stream.read()
    return readme


setuptools.setup(
    name="backend_api",
    version="",
    author="bigcnote",
    author_email="",
    description="",
    long_description=get_package_description(),
    long_description_content_type="text/markdown",
    url="https://github.com/big-c-note/genetic_algo_api",
    packages=setuptools.find_packages(),
    classifiers=[
        "Programming Language :: Python :: 3",
        "Operating System :: OS Independent",
    ],
    python_requires=">=3.8",
)
