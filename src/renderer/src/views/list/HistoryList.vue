<template>
  <PageContainer>
    <a-card class="card" :bordered="false">
      <a-table
        ref="table"
        size="default"
        row-key="key"
        :columns="historyProps.columns"
        :data-source="historyProps.dataSource"
        :alert="true"
      >
      </a-table>
    </a-card>
    <a-drawer
      :title="drawerProps.drawerTitle"
      :closable="false"
      :width="1000"
      :visible="drawerProps.visible"
      :keyboard="true"
      :after-visible-change="queryDetailList"
      @close="handleCancel"
    >
      <a-table
        ref="table"
        size="default"
        row-key="key"
        :columns="drawerProps.drawerColumns"
        :data-source="drawerProps.drawerDataSource"
        :alert="true"
      >
      </a-table>
      <div class="steps-title">日志</div>
      <a-divider />
      <div class="console-log-box">
        <div>
          <pre>
            <span>{{ drawerProps.drawerConsoleLog }}</span>
          </pre>
        </div>
      </div>
    </a-drawer>
    </PageContainer>
</template>

<script setup lang="ts">
const drawerColumns = [
  {
    title: '编号',
    scopedSlots: { customRender: 'serial' }
  },
  {
    title: '功能模块',
    dataIndex: 'Name',
    key: 'Name'
  },
  {
    title: '升级包名称',
    dataIndex: 'Description',
    key: 'Description',
    scopedSlots: { customRender: 'description' }
  },
  {
    title: '操作',
    dataIndex: 'action',
    width: '250px',
    scopedSlots: { customRender: 'action' }
  }
]
const columns = [
  {
    title: '编号',
    scopedSlots: { customRender: 'serial' }
  },
  {
    title: '升级包名称',
    dataIndex: 'Name',
    key: 'Name'
  },
  {
    title: '升级时间',
    dataIndex: 'DateTime'
  },
  {
    title: '操作',
    dataIndex: 'action',
    width: '150px',
    scopedSlots: { customRender: 'action' }
  }
]

const { proxy } = getCurrentInstance()

const historyProps = reactive({
  columns: columns,
  dataSource: []
})

const drawerProps = reactive({
  drawerTitle: '',
  visible: false,
  drawerColumns: drawerColumns,
  drawerDataSource: [],
  drawerConsoleLog: ''
})

function queryDetailList(): void {
  proxy.$message.info('queryDetailList')
}

function handleCancel(): void {
  proxy.$message.info('handleCancel')
}
</script>

<style lang="less" scoped>
.card {
  margin-bottom: 24px;
}

.steps-title {
  font-size: 16px;
  font-weight: bold;
}

.console-log-box {
  background-color: #28384bff;
  min-height: 50px;

  & > div {
    font-size: 1em;
    font-weight: 700;
    padding: 5px 20px;

    & > pre {
      margin-top: 0;
      margin-bottom: 1em;

      & > span {
        color: #e6ebec;
      }
    }
  }
}
</style>
