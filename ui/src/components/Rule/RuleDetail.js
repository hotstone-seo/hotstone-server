import React, { useEffect, useState } from 'react';
import PropTypes from 'prop-types';
import { Link } from 'react-router-dom';
import { Descriptions, message, Divider } from 'antd';
import moment from 'moment';
import { getDataSourcesByIDs } from 'api/datasource';

const formatDate = (dateString) => moment(dateString).format('DD MMM YYYY HH:mm:ss');

function RuleDetail({ rule }) {
  const {
    name,
    url_pattern: urlPattern,
    data_source_ids: dataSourceIDs,
    created_at: createdAt,
    updated_at: updatedAt,
  } = rule;
  const [dataSources, setDataSources] = useState([]);

  useEffect(() => {
    if (dataSourceIDs && dataSourceIDs.length > 0) {
      getDataSourcesByIDs(dataSourceIDs)
        .then((retrievedValue) => {
          setDataSources(retrievedValue);
        })
        .catch((error) => {
          message.error(error.message);
        });
    }
  }, [dataSourceIDs]);

  return (
    <Descriptions>
      <Descriptions.Item key="name" label="Name">{name}</Descriptions.Item>
      <Descriptions.Item key="urlPattern" label="URL Pattern">{urlPattern}</Descriptions.Item>
      <Descriptions.Item key="dataSource" label="Data Source">
        {dataSources.map((dataSource, index) => (
          <span key={dataSource.id}>
            <Link to={`/datasources/${dataSource.id}`}>
              {dataSource.name}
            </Link>
            {index < dataSources.length - 1 && (
              <Divider type="vertical" />
            )}
          </span>
        ))}
      </Descriptions.Item>
      <Descriptions.Item key="createdAt" label="Created at">{formatDate(createdAt)}</Descriptions.Item>
      <Descriptions.Item key="updatedAt" label="Updated at">{formatDate(updatedAt)}</Descriptions.Item>
    </Descriptions>
  );
}

RuleDetail.propTypes = {
  rule: PropTypes.shape({
    name: PropTypes.string,
    url_pattern: PropTypes.string,
    data_source_ids: PropTypes.arrayOf(PropTypes.number),
    created_at: PropTypes.string,
    updated_at: PropTypes.string,
  }).isRequired,
};

export default RuleDetail;
