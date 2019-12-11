import React, { Component } from 'react';
import { UncontrolledDropdown, DropdownItem, DropdownMenu, DropdownToggle, Nav } from 'reactstrap';
import PropTypes from 'prop-types';

import { AppNavbarBrand, AppSidebarToggler } from '@coreui/react';
import logo from '../../assets/img/brand/hotstoneseo_header.png'
import hostoneseo_logo from '../../assets/img/brand/hotstoneseo_logo.png'

const propTypes = {
  children: PropTypes.node,
};

const defaultProps = {};

class DefaultHeader extends Component {
  render() {

    // eslint-disable-next-line
    const { children, ...attributes } = this.props;

    return (
    
      <React.Fragment>
        <AppSidebarToggler className="d-lg-none" display="md" mobile />
        <AppNavbarBrand
          full={{ src: logo, width: 89, height: 21, alt: 'HotStone Logo' }}
          minimized={{ src: hostoneseo_logo, width: 30, height: 30, alt: 'HotStone Logo' }}
        />
        <AppSidebarToggler className="d-md-down-none" display="lg" />
  
        <Nav className="ml-auto" navbar>
         
          <UncontrolledDropdown nav direction="down">
            <DropdownToggle nav>
              <img src={'../../assets/img/avatars/6.jpg'} className="img-avatar" alt="" />
            </DropdownToggle>
            <DropdownMenu right>
     
              <DropdownItem><i className="fa fa-user"></i> Profile</DropdownItem>
              
              <DropdownItem onClick={e => this.props.onLogout(e)}><i className="fa fa-lock"></i> Logout</DropdownItem>
            </DropdownMenu>
          </UncontrolledDropdown>
        </Nav>
       
      </React.Fragment>
    );
  }
}

DefaultHeader.propTypes = propTypes;
DefaultHeader.defaultProps = defaultProps;

export default DefaultHeader;
