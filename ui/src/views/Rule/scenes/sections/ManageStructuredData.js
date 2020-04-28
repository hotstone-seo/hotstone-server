import React, { useState, useEffect } from 'react';
import PropTypes from 'prop-types';
import { message, Space, Button } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import { fetchStructuredDatas, deleteStructuredData } from 'api/structuredData';
import { StructuredDataForm, StructuredDataList } from 'components/StructuredData';

// Aliasing for shorter name
const fetchStructs = fetchStructuredDatas;
const deleteStruct = deleteStructuredData;

function ManageStructuredData({ ruleID }) {
  // NOTE: Structs is the shorthand that we use for Structured Data
  const [focusStruct, setFocusStruct] = useState(null);
  const [structs, setStructs] = useState([]);

  useEffect(() => {
    fetchStructs({ rule_id: ruleID })
      .then((newStructs) => {
        setStructs(newStructs);
      })
      .catch((error) => {
        message.error(error.message);
      });
  }, [ruleID]);

  const refreshStruct = () => {
    fetchStructs({ rule_id: ruleID })
      .then((newStructs) => {
        setStructs(newStructs);
        setFocusStruct(null);
      })
      .catch((error) => {
        message.error(error.message);
      });
  };

  const removeStruct = ({ id: structID }) => {
    deleteStruct(structID)
      .then(() => {
        setStructs(structs.filter((struct) => struct.id !== structID));
      })
      .catch((error) => {
        message.error(error.message);
      });
  };

  return (
    focusStruct ? (
      <StructuredDataForm
        structuredData={focusStruct}
        afterSubmit={refreshStruct}
        onCancel={() => setFocusStruct(null)}
      />
    ) : (
      <Space direction="vertical" style={{ width: '100%' }}>
        <Button
          type="dashed"
          onClick={() => setFocusStruct({ rule_id: ruleID })}
          style={{ width: '100%' }}
        >
          <PlusOutlined />
          Add Structured Data
        </Button>
        <StructuredDataList
          structuredDatas={structs}
          onEdit={(struct) => setFocusStruct(struct)}
          onDelete={removeStruct}
        />
      </Space>
    )
  );
}

ManageStructuredData.propTypes = {
  ruleID: PropTypes.number.isRequired,
};

export default ManageStructuredData;
