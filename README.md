WaitForURL
=========================

This tool checks an URL for content.

Usage
-------------------------
It is possible to configure the URL, the timeout period and a period for checks.
For content validation it is possible to configure a search string.

    usage: ./waitforurl
         -period int
                Timeperiod in seconds for checking port and content (default 30)
          -search string
                Search string
          -timeout int
                Timeout for waiting in seconds for port (default 300)
          -url string
                URL to wait for

    
Example:

    ./waitforurl http://localhost:1007/servlet/ConfigurationServlet


License
------------
Copyright 2014-2019 Matthias Raab.

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
