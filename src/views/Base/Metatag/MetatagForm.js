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
  DropdownItem,
  DropdownMenu,
  DropdownToggle,
  Input,
  InputGroup,
  InputGroupButtonDropdown,
  Label,
  Row,
} from 'reactstrap';

class MetatagForm extends Component {
  constructor(props) {
    super(props);

    this.toggle = this.toggle.bind(this);
    this.toggleFade = this.toggleFade.bind(this);
    this.state = {
      collapse: true,
      fadeIn: true,
      timeout: 300
    };
  }

  toggle() {
    this.setState({ collapse: !this.state.collapse });
  }

  toggleFade() {
    this.setState((prevState) => { return { fadeIn: !prevState }});
  }

  render() {
    return (
      <div className="animated fadeIn">
     
        <Row>
          <Col xs="12" md="12">
            <Card>
              <CardHeader>
                <strong>Add New Meta-Tag</strong>
              </CardHeader>
              <CardBody>
                <Form action="" method="post" encType="multipart/form-data" className="form-horizontal">
                   
                  <FormGroup row>
                    <Col md="3">
                      <Label htmlFor="text-input">Name</Label>
                    </Col>
                    <Col xs="12" md="9">
                      <Input type="text" id="name" name="name" placeholder="Name" />
                      
                    </Col>
                  </FormGroup>
                  <FormGroup row>
                    <Col md="3">
                      <Label htmlFor="text-input">Content</Label>
                    </Col>
                    <Col xs="12" md="9">
                      <Input type="text" id="content" name="content" placeholder="Content" />
                       
                    </Col>
                  </FormGroup>
                   
                  <FormGroup row>
                    <Col md="3">
                      <Label htmlFor="text-input">Default Content</Label>
                    </Col>
                    <Col xs="12" md="9">
                        <Input type="text" id="defaultcontent" name="defaultcontent" placeholder="Default Content" />
                    </Col>
                  </FormGroup>
                  <FormGroup row>
                    <Col md="3">
                      <Label htmlFor="text-input">Rule</Label>
                    </Col>
                    <Col xs="12" md="9">
                        <InputGroup>
                        <InputGroupButtonDropdown addonType="prepend"
                                                  isOpen={this.state.first}
                                                  toggle={() => { this.setState({ first: !this.state.first }); }}>
                          <DropdownToggle caret color="primary">
                            -Choose-
                          </DropdownToggle>
                          <DropdownMenu className={this.state.first ? 'show' : ''}>
                            <DropdownItem>Airport Detail</DropdownItem>                           
                          </DropdownMenu>
                        </InputGroupButtonDropdown>                     
                      </InputGroup>
                    </Col>
                  </FormGroup>
                </Form>
              </CardBody>
              <CardFooter>
                <Button type="submit" size="sm" color="primary"><i className="fa fa-dot-circle-o"></i> Submit</Button>
                <Button type="button" size="sm" color="secondary"><i className="fa fa-ban"></i> Preview</Button>
              </CardFooter>
            </Card>
            
          </Col>
          <Col xs="12" md="6">
             
            
            
             
          </Col>
        </Row>
         
        
         
         
      </div>
    );
  }
}

export default MetatagForm;
