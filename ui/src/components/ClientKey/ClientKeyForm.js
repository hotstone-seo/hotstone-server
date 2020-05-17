import React, { useEffect } from 'react';
import PropTypes from 'prop-types';
import {
  Form, Input, Button, Tooltip,
} from 'antd';
import { QuestionCircleOutlined } from '@ant-design/icons';

function ClientKeyForm({ dataSource, handleSubmit }) {
  const [form] = Form.useForm();

  useEffect(() => {
    form.setFieldsValue(dataSource);
  }, [dataSource, form]);

  return (
    <Form
      form={form}
      onFinish={handleSubmit}
      labelCol={{ span: 6 }}
      wrapperCol={{ span: 14 }}
    >
      <Form.Item name="id" noStyle>
        <Input type="hidden" />
      </Form.Item>
      <Form.Item
        name="name"
        label="Name"
        rules={[{ required: true, message: 'Please input the name of your Client Key' }]}
      >
        <Input data-testid="input-name" placeholder="My Client Key" maxLength="100" />
      </Form.Item>

      <Form.Item
        name="url"
        label="Resource URL"
        rules={[{ required: true, message: 'Please input the Resource URL' }]}
      >
        <Input
          placeholder="http://api.service.com/resource"
          suffix={(
            <Tooltip title="Where the data to fill tags will be retrieved from">
              <QuestionCircleOutlined />
            </Tooltip>
          )}
          maxLength="255"
          data-testid="input-url"
        />
      </Form.Item>

      <Form.Item wrapperCol={{ offset: 6, span: 14 }}>
        <Button data-testid="btn-save" type="primary" htmlType="submit">
          Save
        </Button>
      </Form.Item>
    </Form>
  );
}

ClientKeyForm.defaultProps = {
  dataSource: {},
};

ClientKeyForm.propTypes = {
  dataSource: PropTypes.shape({
    id: PropTypes.number,
    name: PropTypes.string,
    url: PropTypes.string,
  }),

  handleSubmit: PropTypes.func.isRequired,
};

export default ClientKeyForm;
