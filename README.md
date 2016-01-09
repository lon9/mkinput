#mkinput

[![Build Status](https://drone.io/github.com/Rompei/mkinput/status.png)](https://drone.io/github.com/Rompei/mkinput/latest)

mkinput is command to generate input.txt for programs.

Executable [drone.io](https://drone.io/github.com/Rompei/mkinput/files)

##Example

Input [json file](https://github.com/Rompei/mkinput/blob/master/src.json)

```json
{
  "set": 100,           // The number of data sets.
  "sep":"EOF",          // Separator of the data sets.
  "templates":[         // Data set templates.
    {
      "min":1,          // Min value of the data.
      "max":4,          // Max value of the data.
      "minRows": 1,     // Min rows.
      "maxRows": 9,     // Max rows.
      "minCols": 1,     // Min clos.
      "maxCols": 10,    // Max cols.
      "sep": " ",       // Separator of cols.
      "rowSize": true,  // If we write this field, row size is written in front of the data.
      "colSize": false
    }
  ]
}
```

Type command. `mkinput -f <output-filepath> < <source-file>`

```bash
$ mkinput -f "input.txt" < src.json
```

then generate input.txt like [this](https://github.com/Rompei/mkinput/blob/master/input.txt)

```
8
1 2
3 4
3 2
1 3
4 4
1 4
3 4
2 1
EOF
9
2 2
3 4
4 3
3 3
3 3
4 2
2 3
3 2
4 2
EOF

```
