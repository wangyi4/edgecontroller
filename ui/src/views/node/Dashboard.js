import React, { Component } from 'react';
import ApiClient from '../../api/ApiClient';
import { SchemaForm, utils } from 'react-schema-form';
import NodeSchema from '../../components/schema/Node';
import {
  Grid,
  Button
} from '@material-ui/core';

export default class DashboardView extends Component {
  constructor(props) {
    super(props);

    this.state = {
      loaded: false,
      errored: false,
      error: null,
      showErrors: true,
      node: {},
    };
  }

  // GET /nodes/:node_id
  getNode = () => {
    const { nodeID } = this.props;

    ApiClient.get(`/nodes/${nodeID}`)
      .then((resp) => {
        this.setState({
          loaded: true,
          node: resp.data || {},
        })
      })
      .catch((err) => {
        this.setState({
          loaded: true,
          errored: true,
          error: err,
        });
      });
  }

  // PATCH /nodes/:node_id
  updateNode = () => {
    const { nodeID } = this.props;
    const { node } = this.state;

    ApiClient.patch(`/nodes/${nodeID}`, node)
      .then((resp) => {
        this.setState({
          loaded: true,
        })
      })
      .catch((err) => {
        this.setState({
          loaded: true,
          errored: true,
          error: err,
        });
      });
  }

  onModelChange = (key, val) => {
    const { node } = this.state;

    const newNode = node;

    utils.selectOrSet(key, newNode, val);

    this.setState({ node: newNode });
  }

  componentDidMount() {
    this.getNode();
  }

  render() {
    const {
      loaded,
      errored,
      error,
      showErrors,
      node,
    } = this.state;

    if (!loaded) {
      return <React.Fragment>Loading ...</React.Fragment>
    }

    if (errored) {
      return <React.Fragment>{error.toString()}</React.Fragment>
    }

    return (
      <React.Fragment>
        <Grid
          container
          justify="space-between"
          alignItems="flex-end"
          spacing={40}
        >
          <Grid item xs={12}>
            <SchemaForm
              schema={NodeSchema.schema}
              form={NodeSchema.form}
              model={node}
              onModelChange={this.onModelChange}
              showErrors={showErrors}
            />
          </Grid>
          <Grid item xs={12}>
            <Button
              onClick={this.updateNode}
              variant="outlined"
              color="primary"
            >
              Save
            </Button>
          </Grid>
        </Grid>
      </React.Fragment>
    );
  }
};
