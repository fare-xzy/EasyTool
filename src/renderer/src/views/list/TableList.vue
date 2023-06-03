<template>
  <PageContainer>
    <a-card class="card" :bordered="false">
      <div class="table-operator steps-title">
        升级步骤&nbsp;&nbsp;&nbsp;
        <a-popconfirm
          title="是否重新上传？"
          :visible="prepareProps.popvisible"
          placement="right"
          ok-text="Yes"
          cancel-text="No"
          @visible-change="handleVisibleChange"
          @confirm="confirm"
          @cancel="popcancel"
        >
          <a-button type="primary">
            <template #icon><plus-outlined /></template>
            上传升级包
          </a-button>
        </a-popconfirm>
        <a-divider />
        <a-steps
          progress-dot
          class="steps"
          :current="prepareProps.current"
          :status="prepareProps.status"
        >
          <a-step title="上传升级包" />
          <a-step title="解压升级包" />
          <a-step title="检查升级包" />
        </a-steps>
      </div>
    </a-card>

    <a-card class="card" :bordered="false">
      <div class="steps-title">
        升级包详情
        <a-button
          class="update-button"
          type="primary"
          :loading="detailProps.isLoading"
          @click="updateNow"
          >立即升级</a-button
        >
      </div>
      <a-divider />
      <a-table
        ref="table"
        size="default"
        row-key="key"
        :columns="detailProps.columns"
        :data-source="detailProps.dataSource"
        :alert="true"
      >
      </a-table>

      <!-- <create-form
        ref="createModal"
        :visible="visible"
        :loading="confirmLoading"
        :model="mdl"
        @cancel="handleCancel"
        @ok="handleOk"
      /> -->
    </a-card>
    <a-card class="card" :bordered="false">
      <div class="steps-title">升级步骤&nbsp;&nbsp;&nbsp; {{ updateProps.stepTitle }}</div>
      <a-divider />
      <a-steps
        progress-dot
        class="steps"
        :current="updateProps.current"
        :status="updateProps.status"
      >
        <a-step title="执行备份" />
        <a-step title="升级" />
        <a-step title="完成" />
      </a-steps>
    </a-card>
    <a-card class="card" :bordered="false">
      <div class="steps-title">升级日志</div>
      <a-divider />
      <div class="console-log-box">
        <div>
          <pre>
            <span>{{ updateProps.consoleLog }}</span>
          </pre>
        </div>
      </div>
    </a-card>
  </PageContainer>
</template>

<script setup lang="ts">
const { proxy } = getCurrentInstance()
const columns = [
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

const prepareProps = reactive({
  popvisible: false,
  current: 0,
  status: 'wait'
})
const detailProps = reactive({
  isLoading: false,
  columns: columns,
  dataSource: []
})
const updateProps = reactive({
  current: 0,
  status: 'wait',
  stepTitle: '',
  consoleLog: ''
})

function confirm(): void {
  proxy.$message.info('confirm')
}
function popcancel(): void {
  prepareContents.popvisible = false
}
function handleVisibleChange(params): void {
  proxy.$message.info('handleVisibleChange' + params)
}
function updateNow(): void {
  proxy.$message.info('updateNow')
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

.update-button {
  margin-left: 20px;
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
