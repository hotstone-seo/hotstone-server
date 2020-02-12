import React from 'react';
import { Form, Input, Select } from 'antd';
import locales from 'locales';

const { Option } = Select;

function MetaForm({ tag }) {
  const [form] = Form.useForm();
  if (tag) {
    form.setFieldsValue(tag);
  }
  form.setFieldsValue({ type: 'meta' });

  return (
    <Form form={form}>
      <Form.Item label="Locale" name="locale">
        <Select>
          {locales.map(locale => (
            <Option value={locale}>{locale}</Option>
          ))}
        </Select>
      </Form.Item>
      <Form.Item label="Name" name={['attributes', 'name']}>
        <Input />
      </Form.Item>
      <Form.Item label="Content" name={['attributes', 'content']}>
        <Input />
      </Form.Item>
    </Form>
  );
}

export default MetaForm;