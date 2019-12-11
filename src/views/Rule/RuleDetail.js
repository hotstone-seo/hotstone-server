import React, { Component } from 'react';
import {
    Button,
    Card,
    CardBody,
    CardFooter,
    CardHeader,
    Col,
    Form,
    FormGroup,
    Label,
    NavLink,
    Pagination,
    PaginationItem,
    PaginationLink,
    Row,
    Table,
} from 'reactstrap';

import PropTypes from 'prop-types';


class RuleDetail extends Component {
    constructor(props) {
        super(props);

        this.toggle = this.toggle.bind(this);
        this.toggleFade = this.toggleFade.bind(this);
        this.state = {
            collapse: true,
            fadeIn: true,
            timeout: 300
        };
        this.handleBack = this.handleBack.bind(this);
        this.handleEditCanonical = this.handleEditCanonical.bind(this);

        this.handleAddNewCanonical = this.handleAddNewCanonical.bind(this);
        this.handleAddNewMeta = this.handleAddNewMeta.bind(this);
        this.handleAddNewScript = this.handleAddNewScript.bind(this);
        this.handleAddNewTitle = this.handleAddNewTitle.bind(this);
    }

    toggle() {
        this.setState({ collapse: !this.state.collapse });
    }

    toggleFade() {
        this.setState((prevState) => { return { fadeIn: !prevState } });
    }
    handleBack() {
        const { history } = this.props;
        history.push('/rule');
    }
    handleEditCanonical() {
        const { history } = this.props;
        history.push('/canonicalEditForm');
    }
    handleAddNewCanonical(){
        const { history } = this.props;
        history.push('/canonicalForm');
    }
    handleAddNewMeta(){
        const { history } = this.props;
        history.push('/metatagForm');
    }
    handleAddNewScript(){
        const { history } = this.props;
        history.push('/scripttagForm');
    }
    handleAddNewTitle(){
        const { history } = this.props;
        history.push('/titletagForm');
    }

    render() {
        const ruleID = this.props.match.params.id;

        return (
            <div className="animated fadeIn">
                <Row>
                    <Col xs="12" md="9" lg="6">
                        <Card>
                            <CardHeader>
                                <strong>Detail Rule ID {ruleID}</strong>
                            </CardHeader>
                            <CardBody>
                                <Form action="" method="post" encType="multipart/form-data" className="form-horizontal">
                                    <FormGroup row>
                                        <Col md="3">
                                            <Label htmlFor="text-input">Name</Label>
                                        </Col>
                                        <Col xs="12" md="9">
                                            Airport Detail
                                        </Col>
                                    </FormGroup>
                                    <FormGroup row>
                                        <Col md="3">
                                            <Label htmlFor="text-input">URL Pattern</Label>
                                        </Col>
                                        <Col xs="12" md="9">
                                            http://xxx
                                        </Col>
                                    </FormGroup>
                                    <FormGroup row>
                                        <Col md="3">
                                            <Label htmlFor="textarea-input">Exclusion</Label>
                                        </Col>
                                        <Col xs="12" md="9">
                                            -
                                        </Col>
                                    </FormGroup>
                                    <FormGroup row>
                                        <Col md="3">
                                            <Label htmlFor="text-input">Data Source</Label>
                                        </Col>
                                        <Col xs="12" md="9">
                                            Airport
                                        </Col>
                                    </FormGroup>
                                </Form>
                            </CardBody>
                            <CardFooter>

                                <Button type="button" size="md" color="secondary" onClick={this.handleBack}> Back</Button>
                            </CardFooter>
                        </Card>
                    </Col>
                </Row>

                <Row>
                    <Col>
                        <Card>
                            <CardHeader>
                                <i className="fa fa-align-justify"></i> 
                            </CardHeader>
                            <CardBody>
                                <div style={{ marginBottom: '.5rem' }}>
                                    <Button color="primary" onClick={this.handleAddNewCanonical}>Add New Canonical</Button>
                                    <Button color="primary" onClick={this.handleAddNewMeta}>Add New Meta-Tag</Button>
                                    <Button color="primary" onClick={this.handleAddNewScript}>Add New Script Tag</Button>
                                    <Button color="primary" onClick={this.handleAddNewTitle}>Add New Title-Tag</Button>
                                </div>
                                <Table responsive bordered>
                                    <thead>
                                        <tr>
                                            <th>Type</th>
                                            <th>Attribute</th>
                                            <th>Value</th>
                                           
                                            <th>Action</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        <tr>
<<<<<<< HEAD
                                            <td>Canonical</td>
                                            <td>xxx</td>
                                            <td>http://tiket.com/asad</td>
=======
                                            <td>http://hotstone-seo/asad</td>
>>>>>>> 66f6786f1dfff735101127dfef630879fcff0a37

                                            
                                            <td>
                                                <NavLink href="#" onClick={this.handleEdit}>Edit</NavLink>
                                            </td>
                                        </tr>
                                    </tbody>
                                </Table>
                                <nav>
                                    <Pagination>
                                        <PaginationItem><PaginationLink previous tag="button">Prev</PaginationLink></PaginationItem>
                                        <PaginationItem active>
                                            <PaginationLink tag="button">1</PaginationLink>
                                        </PaginationItem>

                                        <PaginationItem><PaginationLink next tag="button">Next</PaginationLink></PaginationItem>
                                    </Pagination>
                                </nav>
                            </CardBody>
                        </Card>
                    </Col>
                </Row>
                
                 
            </div>
        );
    }
}
RuleDetail.propTypes = {
    match: PropTypes.shape({
        path: PropTypes.string,
    }).isRequired,
};
export default RuleDetail;