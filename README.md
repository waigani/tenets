# Tenets Repository

This repository hosts published Tenets that can be imported as Tenet bundles or individually in .lingo files. The file structure follows the following format:

```
<owner>/
  <bundle>/
    <tenet>/
      .lingo
      [example files...]
```

The owner corresponds to a CodeLingo user account name (you can sign up at codelingo.io). To land code you must have permissions to push to the owner directory.

Tenet bundles are collections of Tenets grouped by theme: usually language, stack or company. Each Tenet in the bundle is within it's own folder, it is important that the name of the folder match the name of the Tenet in the .lingo file.

Tenets are activated by Flows. To test an individual Tenet, cd into it's folder and run a Flow which activates any bots specified in the Tenet. A common bot is codelingo/review which can be activated by running the review Flow. To run the review flow, run the following command within this repository: `lingo review -i`.
