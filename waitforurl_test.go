/*
 * Copyright (c) 2019.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestWithoutServer(t *testing.T) {
	testconfig := &Config{}
	testconfig.runtime = 10
	testconfig.timeout = 50
	testconfig.period = 5

	testconfig.urlString = "http://localhost/hello"

	err := testconfig.CheckForContent()

	if err == nil {
		t.Fatalf("Error ...")
	}
}

func TestRequestWithoutSearchString(t *testing.T) {
	testconfig := &Config{}
	testconfig.runtime = 10
	testconfig.timeout = 100
	testconfig.period = 2

	testrun := 0

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/hello" {
			http.Error(w, "404 not found.", http.StatusNotFound)
			return
		}

		if r.Method != "GET" {
			http.Error(w, "Method is not supported.", http.StatusNotFound)
			return
		}

		if testrun < 3 {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintf(w, "STARTING")
			testrun = testrun + 1
		} else {
			fmt.Fprintf(w, "RUNNING")
		}
	}))

	defer ts.Close()
	testconfig.urlString = ts.URL + "/hello"

	err := testconfig.CheckForContent()

	if err != nil {
		t.Fatalf("Error ...")
	}

}

func TestRequestWithSearchString(t *testing.T) {
	testconfig := &Config{}
	testconfig.runtime = 10
	testconfig.timeout = 100
	testconfig.period = 2
	testconfig.searchString = "RUNNING"
	testrun := 0

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/hello" {
			http.Error(w, "404 not found.", http.StatusNotFound)
			return
		}

		if r.Method != "GET" {
			http.Error(w, "Method is not supported.", http.StatusNotFound)
			return
		}

		if testrun < 3 {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintf(w, "STARTING")
			testrun = testrun + 1
		}

		if testrun < 4 {
			fmt.Fprintf(w, "WAITING")
			testrun = testrun + 1
		} else {
			fmt.Fprintf(w, "RUNNING")
		}
	}))

	defer ts.Close()
	testconfig.urlString = ts.URL + "/hello"

	err := testconfig.CheckForContent()

	if err != nil {
		t.Fatalf("Error ...")
	}

}

func TestRequestWithOtherSearchString(t *testing.T) {
	testconfig := &Config{}
	testconfig.runtime = 10
	testconfig.timeout = 100
	testconfig.period = 2
	testconfig.searchString = "OTHER"
	testrun := 0

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/hello" {
			http.Error(w, "404 not found.", http.StatusNotFound)
			return
		}

		if r.Method != "GET" {
			http.Error(w, "Method is not supported.", http.StatusNotFound)
			return
		}

		if testrun < 3 {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintf(w, "STARTING")
			testrun = testrun + 1
		}

		if testrun < 4 {
			fmt.Fprintf(w, "WAITING")
			testrun = testrun + 1
		} else {
			fmt.Fprintf(w, "RUNNING")
		}
	}))

	defer ts.Close()
	testconfig.urlString = ts.URL + "/hello"

	err := testconfig.CheckForContent()

	if err == nil {
		t.Fatalf("Error ...")
	}

}
