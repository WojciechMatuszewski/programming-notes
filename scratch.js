require("./module");
delete require.cache[require.resolve("./module")];
require("./module");
