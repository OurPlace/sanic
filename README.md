# Sanic

This application performs speedtests using the speedtest-cli, and then stores the results in a database. 

## Configuration 

This application is configured with a json file that looks like this: 

``` json
{
    "dsn":"dsn",
    "loop":"10m"
}
```

## Speedtest CLI

**Tested Version**: 1.0.0.2

The speedtest CLI _must_ be present in the path. To verify that you have it
installed correctly, you may run the command `speedtest --version` and expect to
see an output similar to this:

```
Speedtest by Ookla 1.0.0.2 (5ae238b) Darwin 19.3.0 x86_64

The official command line client for testing the speed and performance
of your internet connection.
```

Instructions on how to install the CLI can be found
[here](https://www.speedtest.net/apps/cli) and a blog post by speedtest.net
talking about this tool can be found
[here](https://www.speedtest.net/insights/blog/introducing-speedtest-cli/).

## License

```
    Copyright 2020 deadly.surgery

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use these files except in compliance with the License.
    You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
```