<template>
  <div class="createPost-container">
    <el-form ref="postForm" :model="postForm" :rules="rules" class="form-container">


      <div class="createPost-main-container">
        <el-row>
          <el-col :span="6">
            <el-form-item style="margin-bottom: 40px;" prop="name">
              <MDinput v-model="postForm.name" :maxlength="20" name="name" required>
                标题
              </MDinput>
            </el-form-item>
          </el-col>

          <el-col :span="6">
            <el-form-item style="margin-bottom: 40px;" label="标签">
              <el-select v-model="postForm.tags" value-key="id" multiple filterable clearable placeholder="请选择">
                <el-option v-for="item in tags" :key="item.id" :label="item.name" :value="item">
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>

          <el-col :span="6">
            <el-button v-loading="loading" style="margin-left: 10px;" type="success" @click="submitForm">
              发布
            </el-button>
            <el-button v-loading="loading" type="warning" @click="draftForm">
              草稿
            </el-button>
          </el-col>
        </el-row>

        <el-form-item style="margin-bottom: 40px;" prop="desc">
          <MDinput v-model="postForm.desc" :maxlength="100" name="name" required>
            描述
          </MDinput>
        </el-form-item>

        <el-form-item prop="content">
          <markdown-editor ref="editor" v-model="postForm.content"
            :options="{ hideModeSwitch: true, previewStyle: 'tab' }" height="600px" />
          <span v-show="contentShortLength" class="word-counter">{{ contentShortLength }}</span>
        </el-form-item>
      </div>
    </el-form>
  </div>
</template>

<script>
import MarkdownEditor from '@/components/MarkdownEditor'
import MDinput from '@/components/MDinput'
import { getArticle, createArticle, updateArticle } from '@/api/article'
import { listTag } from '@/api/tag'

const defaultForm = {
  id: undefined,
  name: '', // 文章题目
  desc: '', // 文章描述
  content: '', // 文章内容
  state: undefined,
  tags: [],
  create_time: null,
  update_time: null
}

export default {
  name: 'ArticleDetail',
  components: { MarkdownEditor, MDinput },
  props: {
    isEdit: {
      type: Boolean,
      default: false
    }
  },
  filters: {
    statusTypeFilter (status) {
      const statusMap = {
        1: 'publish',
        2: 'draft',
        3: 'deleted'
      }
      return statusMap[status]
    },
  },
  data () {
    const validateRequire = (rule, value, callback) => {
      if (value === '') {
        this.$message({
          message: rule.field + '为必传项',
          type: 'error'
        })
        callback(new Error(rule.field + '为必传项'))
      } else {
        callback()
      }
    }
    return {
      postForm: Object.assign({}, defaultForm),
      loading: false,
      userListOptions: [],
      tags: [],
      rules: {
        //cover_image_url: [{ validator: validateRequire }],
        name: [{ validator: validateRequire }],
        content: [{ validator: validateRequire }],
      },
      tempRoute: {}
    }
  },
  computed: {
    contentShortLength () {
      return this.postForm.content.length
    },
  },
  created () {
    if (this.isEdit) {
      const id = this.$route.params && this.$route.params.id
      this.fetchData(id)
      this.getTags()
    }

    // Why need to make a copy of this.$route here?
    // Because if you enter this page and quickly switch tag, may be in the execution of the setTagsViewTitle function, this.$route is no longer pointing to the current page
    // https://github.com/PanJiaChen/vue-element-admin/issues/1221
    this.tempRoute = Object.assign({}, this.$route)
  },
  methods: {
    fetchData (id) {
      getArticle(id).then(response => {
        this.postForm = response.data
        // set tagsview title
        this.setTagsViewTitle()

        // set page title
        this.setPageTitle()
      }).catch(err => {
        console.log(err)
      })
    },
    getTags () {
      listTag({ state: 1 }).then(response => {
        this.tags = response.data
      })
    },
    setTagsViewTitle () {
      const title = '编辑文章'
      const route = Object.assign({}, this.tempRoute, { title: `${title}-${this.postForm.id}` })
      //this.$store.dispatch('tagsView/updateVisitedView', route)
    },
    setPageTitle () {
      const title = '编辑文章'
      document.title = `${title} - ${this.postForm.id}`
    },
    submitForm () {
      console.log(this.postForm)
      this.$refs.postForm.validate(valid => {
        if (valid) {
          this.postForm.state = 1
          if (this.isEdit) {
            updateArticle(this.postForm).then(response => {
              this.loading = true
              this.$notify({
                title: '成功',
                message: '修改文章成功',
                type: 'success',
                duration: 2000
              })
              this.loading = false
            })
          } else {
            createArticle(this.postForm).then(response => {
              this.loading = true
              this.$notify({
                title: '成功',
                message: '创建文章成功',
                type: 'success',
                duration: 2000
              })
              this.loading = false
            })
          }
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    draftForm () {
      if (this.postForm.content.length === 0 || this.postForm.name.length === 0) {
        this.$message({
          message: '请填写必要的标题和内容',
          type: 'warning'
        })
        return
      }
      this.postForm.state = 2
      if (this.isEdit) {
        updateArticle(this.postForm).then(response => {
          this.$message({
            message: '保存成功',
            type: 'success',
            showClose: true,
            duration: 1000
          })
        })
      } else {
        createArticle(this.postForm).then(response => {
          this.$message({
            message: '创建成功',
            type: 'success',
            showClose: true,
            duration: 1000
          })
        })
      }
    },
    getTagList (query) {
      listTag(query).then(response => {
        if (!response.data.items) return
        this.userListOptions = response.data.items.map(v => v.name)
      })
    }
  }
}
</script>

<style lang="scss" scoped>
@import "~@/styles/mixin.scss";

.createPost-container {
  position: relative;

  .createPost-main-container {
    padding: 40px 45px 20px 50px;

    .postInfo-container {
      position: relative;
      @include clearfix;
      margin-bottom: 10px;

      .postInfo-container-item {
        float: left;
      }
    }
  }

  .word-counter {
    width: 40px;
    position: absolute;
    right: 10px;
    top: 0px;
  }
}

.article-textarea ::v-deep {
  textarea {
    padding-right: 40px;
    resize: none;
    border: none;
    border-radius: 0px;
    border-bottom: 1px solid #bfcbd9;
  }
}
</style>
