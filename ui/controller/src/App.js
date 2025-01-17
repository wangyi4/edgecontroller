// Copyright 2019 Intel Corporation and Smart-Edge.com, Inc. All rights reserved
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import React, { Component } from 'react';
import { BrowserRouter } from 'react-router-dom';
import { MuiThemeProvider } from '@material-ui/core/styles';
import './App.css';
import Routes from './routes';
import MuiTheme from './MuiTheme';
import '@material-ui/icons';
import 'typeface-lato';
import 'typeface-roboto';
import { SnackbarProvider } from 'notistack';
import CssBaseline from '@material-ui/core/CssBaseline';
import ErrorBoundary from './components/common/ErrorBoundary';
import OrchestrationContext, {
  initialOrchestrationValue,
} from './context/orchestrationContext';

class App extends Component {
  render() {
    return (
      <React.Fragment>
        <CssBaseline />
        <SnackbarProvider
          maxSnack={3}
          anchorOrigin={{
            vertical: 'top',
            horizontal: 'right',
          }}
        >
          <div>
            <ErrorBoundary>
              <MuiThemeProvider theme={MuiTheme}>
                <BrowserRouter>
                  <OrchestrationContext.Provider
                    value={initialOrchestrationValue}
                  >
                    <Routes />
                  </OrchestrationContext.Provider>
                </BrowserRouter>
              </MuiThemeProvider>
            </ErrorBoundary>
          </div>
        </SnackbarProvider>
      </React.Fragment>
    );
  }
}

export default App;
