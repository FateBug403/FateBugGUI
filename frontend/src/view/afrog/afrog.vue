<template>
  <div>
    <vxe-grid ref="xGrid" v-bind="gridOptions" v-on="gridEvent"></vxe-grid>
    <el-dialog v-model="DisableDialog.PocDetailsDialog" :fullscreen=true title="POC详情" destroy-on-close>
      <poc-details :fatherPocString="PocString"></poc-details>
    </el-dialog>
  </div>
</template>

<script lang="ts">
export default { // 用户保持keep-alive
  name: 'afrog'
}
</script>

<script lang="ts" setup>
import { reactive, ref } from 'vue'
import { VxeGridInstance, VxeGridListeners, VxeGridProps } from 'vxe-table'
import {
  GetAfrogList,
  ReadAfrogPoc
} from '../../../wailsjs/go/main/App'
import PocDetails from "./pocDetails.vue";

// 设置poc探测目标
let PocString = reactive({
  context:''
})

// 弹窗显示控制
const DisableDialog = reactive({
  PocDetailsDialog: false,
})


// 设置POC表格
const xGrid = ref<VxeGridInstance>()
const options = {
  border: true,
  showHeaderOverflow: true,
  showOverflow: true,
  keepSource: true,
  autoResize: true,
  id: 'full_edit_1',
  height: 1000,
  rowConfig: {
    keyField: 'id',
    isHover: true
  },
  columnConfig: {
    resizable: true
  },
  editConfig: {
    trigger: 'manual',
    mode: 'row',
    showStatus: true
  },
  customConfig: {
    storage: true,
    checkMethod({ column }) {
      if (['nickname', 'role'].includes(column.field)) {
        return false
      }
      return true
    }
  },
  sortConfig: {
    // trigger: 'cell',
    // remote: true
  },
  filterConfig: {
    remote: false
  },
  pagerConfig: {
    pageSize: 10,
    pageSizes: [5, 10, 15, 20, 50, 100, 200, 500, 1000]
  },
  formConfig: {
    data: {
      tech: '',
      sex: '',
      host: '',
      title: '',
      url: '',
      finger: ''
    },
    items: [
      { field: 'keyword', title: '关键词', span: 8, itemRender: { name: '$input', props: { placeholder: '请输入关键词' }}},
      { field: 'description', title: '描述', span: 8, itemRender: { name: '$input', props: { placeholder: '请输入描述' }}},
      { span: 24, align: 'center', collapseNode: true, itemRender: { name: '$buttons', children: [{ props: { type: 'submit', content: '搜索', status: 'primary' }}, { props: { type: 'reset', content: '清空' }}] }}
      // { slots: { default: 'operate_item' }}
    ]
  },
  toolbarConfig: {
    custom: true, // 显示自定义列按钮
    slots: {
      buttons: 'toolbar_buttons',
      // tools: 'toolbar_tools'
    },
    refresh: true, // 显示刷新按钮
    zoom: true, // 显示全屏按钮

  },
  proxyConfig: {
    seq: true, // 启用动态序号代理，每一页的序号会根据当前页数变化
    sort: true, // 启用排序代理，当点击排序时会自动触发 query 行为
    filter: false, // 启用筛选代理，当点击筛选时会自动触发 query 行为
    form: true, // 启用表单代理，当点击表单提交按钮时会自动触发 reload 行为
    // autoLoad: false, // 自动请求数据
    props: {
      // 对应响应结果 Promise<{ result: [], page: { total: 100 } }>
      result: 'date', // 配置响应结果列表字段
      total: 'page.total' // 配置响应结果总页数字段
    },
    // 只接收Promise，具体实现自由发挥
    ajax: {
      // 当点击工具栏查询按钮或者手动提交指令 query或reload 时会被触发
      query: ({ page, sorts, filters, form }) => {
        const queryParams: any = Object.assign({}, form) // 获取查询参数
        return GetAfrogList(queryParams)
      },
    }
  },
  columns: [
    { type: 'checkbox', title: '状态', width: 15 },
    { field: 'index', title: '序号', sortable: true, width: 20 },
    { field: 'id', title: '标识', sortable: true },
    { field: 'name', title: '名字', sortable: true },
    { field: 'severity', title: '等级', sortable: true,
      filters: [
        { label: 'info', value: 'info' },
        { label: 'medium', value: 'medium' },
        { label: 'high', value: 'high' },
        { label: 'critical', value: 'critical' }
      ],
      filterMethod({ value, option, cellValue, row, column }) {
        return value === row.severity
      }
    },
    { field: 'description', title: '描述', sortable: true },
  ],
}

interface RowVO {
  [key: string]: any
}
const gridOptions = reactive<VxeGridProps<RowVO>>(options)

const gridEvent: VxeGridListeners<RowVO> = {
  async proxyQuery() {
    // const list = await getSearchRecordList({page: 1, pageSize: 10})
    console.log('数据代理查询事件')
  },
  proxyDelete() {
    console.log('数据代理删除事件')
  },
  proxySave() {
    console.log('数据代理保存事件')
  },
  async cellDblclick({row}) {
    //获取指定文件内容
    console.log(row.id)
    const result = await ReadAfrogPoc(row.id)
    if (result.code === 200) {
      // console.log(result)
      PocString.context = result.date
      DisableDialog.PocDetailsDialog = true
    }
  }
}
</script>
<style>
</style>
