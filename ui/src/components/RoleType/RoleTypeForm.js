import React, { useEffect, useState } from 'react';
import PropTypes from 'prop-types';
import {
  Form, Input, Button, Checkbox,
} from 'antd';
import { createRoleType } from 'api/roleType';
import { fetchModules } from 'api/module';

const CheckboxGroup = Checkbox.Group;
const defaultCheckedList = [];

function RoleTypeForm({ roleType, handleSubmit }) {
  const [form] = Form.useForm();
  const [arrayCheck, setArrayCheck] = useState([]);
  const [checkedList, setCheckedList] = useState({
    checkedList: defaultCheckedList,
    indeterminate: true,
    checkAll: false,
    checkboxArray: [],
  });
  const [plainOptions, setPlainOptions] = useState([]);
  const [modulesList, setModuleList] = useState([]);

  async function fetchModulesList() {
    try {
      const modulesAPI = await fetchModules();
      let mn = null;
      let noModule = 0;
      const plainOptionsTemp = [];
      const arrMenu = [];
      if (modulesAPI) {
        Object.keys(modulesAPI).forEach((key) => {
          mn = modulesAPI[key];

          // for label in UI
          plainOptionsTemp[noModule] = mn.label;
          noModule += 1;

          // for checking before process to API create Role
          const tempMenu = [];
          tempMenu.label = mn.label;
          tempMenu.name = mn.name;
          tempMenu.id = mn.id;
          arrMenu.push(tempMenu);
        });
        setPlainOptions(plainOptionsTemp);
        setModuleList(arrMenu);
      }
    } catch (error) {
      console.log(error, 'error getModules');
    }
  }
  useEffect(() => {
    form.setFieldsValue(roleType);
    fetchModulesList();
  }, [roleType, form]);

  const handleonChange = (checkedListNew) => {
    setArrayCheck({ ...arrayCheck, checkedListNew });
    if (checkedListNew) {
      setCheckedList({
        ...checkedList,
        checkedList: checkedListNew,
        indeterminate: !!checkedListNew.length && checkedListNew.length < plainOptions.length,
        checkAll: checkedListNew.length === plainOptions.length,
      });
    }
  };

  const onCheckAllChange = (e) => {
    setCheckedList({
      checkedList: e.target.checked ? plainOptions : [],
      indeterminate: false,
      checkAll: e.target.checked,
    });
  };
  const onSubmit = createRoleType;

  const onFinish = (values) => {
    const checkListFinal = checkedList.checkedList;
    const modulesValue = [];
    Object.keys(modulesList).forEach((keyList) => {
      Object.keys(checkListFinal).forEach((key) => {
        if (checkListFinal[key] === modulesList[keyList].label) {
          const arr = {};
          arr.name = modulesList[keyList].name;
          modulesValue.push(arr);
        }
      });
    });
    values.modules = modulesValue;
    onSubmit(values)
      .then((response) => {
        handleSubmit(response);
      })
      .catch(() => {
      });
  };

  return (
    <Form
      form={form}
      onFinish={onFinish}
      labelCol={{ span: 6 }}
      wrapperCol={{ span: 14 }}
    >
      <Form.Item name="id" noStyle>
        <Input type="hidden" />
      </Form.Item>

      <Form.Item
        name="name"
        label="Role Name"
        rules={[{ required: true, message: 'Please input the role' }]}
      >
        <Input data-testid="input-role-type" placeholder="Role Name" maxLength="200" />
      </Form.Item>

      <Form.Item
        name="modules"
        label="Role Privilege"
      >
        <div>
          <div className="site-checkbox-all-wrapper">
            <Checkbox
              indeterminate={checkedList.indeterminate}
              onChange={onCheckAllChange}
              checked={checkedList.checkAll}
            >
              Check all
            </Checkbox>
          </div>
          <br />
          <CheckboxGroup
            options={plainOptions}
            value={checkedList.checkedList}
            onChange={handleonChange}
            name="modules"
          />
        </div>
      </Form.Item>

      <Form.Item
        wrapperCol={{ offset: 6, span: 14 }}
      >
        <Button data-testid="btn-save" type="primary" htmlType="submit">
          Save
        </Button>
      </Form.Item>
    </Form>
  );
}

RoleTypeForm.defaultProps = {
  roleType: {},
};

RoleTypeForm.propTypes = {
  roleType: PropTypes.shape({
    id: PropTypes.number,
    name: PropTypes.string,
  }),

  handleSubmit: PropTypes.func.isRequired,
};

export default RoleTypeForm;
