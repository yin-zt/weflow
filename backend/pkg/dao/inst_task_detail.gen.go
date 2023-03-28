// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"wego2023/weflow/pkg/model"
)

func newInstTaskDetail(db *gorm.DB, opts ...gen.DOOption) instTaskDetail {
	_instTaskDetail := instTaskDetail{}

	_instTaskDetail.instTaskDetailDo.UseDB(db, opts...)
	_instTaskDetail.instTaskDetailDo.UseModel(&model.InstTaskDetail{})

	tableName := _instTaskDetail.instTaskDetailDo.TableName()
	_instTaskDetail.ALL = field.NewAsterisk(tableName)
	_instTaskDetail.ID = field.NewInt64(tableName, "id")
	_instTaskDetail.InstTaskID = field.NewString(tableName, "inst_task_id")
	_instTaskDetail.FlowDefID = field.NewString(tableName, "flow_def_id")
	_instTaskDetail.ModelID = field.NewString(tableName, "model_id")
	_instTaskDetail.FormDefID = field.NewString(tableName, "form_def_id")
	_instTaskDetail.VersionID = field.NewInt64(tableName, "version_id")
	_instTaskDetail.VersionNum = field.NewString(tableName, "version_num")
	_instTaskDetail.CreateSrc = field.NewInt32(tableName, "create_src")
	_instTaskDetail.TaskName = field.NewString(tableName, "task_name")
	_instTaskDetail.Status = field.NewInt32(tableName, "status")
	_instTaskDetail.Remark = field.NewString(tableName, "remark")
	_instTaskDetail.CreateTime = field.NewTime(tableName, "create_time")
	_instTaskDetail.CreateUser = field.NewString(tableName, "create_user")
	_instTaskDetail.UpdateTime = field.NewTime(tableName, "update_time")
	_instTaskDetail.UpdateUser = field.NewString(tableName, "update_user")
	_instTaskDetail.StartTime = field.NewTime(tableName, "start_time")
	_instTaskDetail.EndTime = field.NewTime(tableName, "end_time")
	_instTaskDetail.SourceID = field.NewString(tableName, "source_id")

	_instTaskDetail.fillFieldMap()

	return _instTaskDetail
}

type instTaskDetail struct {
	instTaskDetailDo instTaskDetailDo

	ALL        field.Asterisk
	ID         field.Int64  // 唯一id
	InstTaskID field.String // 实例任务id
	FlowDefID  field.String // 流程定义id
	ModelID    field.String // 模板id
	FormDefID  field.String // 表单定义id
	VersionID  field.Int64  // 版本id
	VersionNum field.String // 版本号
	CreateSrc  field.Int32  // 创建来源【1：系统发起；2：API发起】
	TaskName   field.String // 实例任务名称
	Status     field.Int32  // 任务状态【1：创建中；2：进行中； 3：终止； 4：完成； 5：挂起；6：草稿】
	Remark     field.String // 描述
	CreateTime field.Time   // 创建时间
	CreateUser field.String // 创建人
	UpdateTime field.Time   // 更新时间
	UpdateUser field.String // 更新人
	StartTime  field.Time   // 发起时间
	EndTime    field.Time   // 结束时间
	SourceID   field.String // 来源id，界面发起人或者api对应的应用

	fieldMap map[string]field.Expr
}

func (i instTaskDetail) Table(newTableName string) *instTaskDetail {
	i.instTaskDetailDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i instTaskDetail) As(alias string) *instTaskDetail {
	i.instTaskDetailDo.DO = *(i.instTaskDetailDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *instTaskDetail) updateTableName(table string) *instTaskDetail {
	i.ALL = field.NewAsterisk(table)
	i.ID = field.NewInt64(table, "id")
	i.InstTaskID = field.NewString(table, "inst_task_id")
	i.FlowDefID = field.NewString(table, "flow_def_id")
	i.ModelID = field.NewString(table, "model_id")
	i.FormDefID = field.NewString(table, "form_def_id")
	i.VersionID = field.NewInt64(table, "version_id")
	i.VersionNum = field.NewString(table, "version_num")
	i.CreateSrc = field.NewInt32(table, "create_src")
	i.TaskName = field.NewString(table, "task_name")
	i.Status = field.NewInt32(table, "status")
	i.Remark = field.NewString(table, "remark")
	i.CreateTime = field.NewTime(table, "create_time")
	i.CreateUser = field.NewString(table, "create_user")
	i.UpdateTime = field.NewTime(table, "update_time")
	i.UpdateUser = field.NewString(table, "update_user")
	i.StartTime = field.NewTime(table, "start_time")
	i.EndTime = field.NewTime(table, "end_time")
	i.SourceID = field.NewString(table, "source_id")

	i.fillFieldMap()

	return i
}

func (i *instTaskDetail) WithContext(ctx context.Context) *instTaskDetailDo {
	return i.instTaskDetailDo.WithContext(ctx)
}

func (i instTaskDetail) TableName() string { return i.instTaskDetailDo.TableName() }

func (i instTaskDetail) Alias() string { return i.instTaskDetailDo.Alias() }

func (i *instTaskDetail) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *instTaskDetail) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 18)
	i.fieldMap["id"] = i.ID
	i.fieldMap["inst_task_id"] = i.InstTaskID
	i.fieldMap["flow_def_id"] = i.FlowDefID
	i.fieldMap["model_id"] = i.ModelID
	i.fieldMap["form_def_id"] = i.FormDefID
	i.fieldMap["version_id"] = i.VersionID
	i.fieldMap["version_num"] = i.VersionNum
	i.fieldMap["create_src"] = i.CreateSrc
	i.fieldMap["task_name"] = i.TaskName
	i.fieldMap["status"] = i.Status
	i.fieldMap["remark"] = i.Remark
	i.fieldMap["create_time"] = i.CreateTime
	i.fieldMap["create_user"] = i.CreateUser
	i.fieldMap["update_time"] = i.UpdateTime
	i.fieldMap["update_user"] = i.UpdateUser
	i.fieldMap["start_time"] = i.StartTime
	i.fieldMap["end_time"] = i.EndTime
	i.fieldMap["source_id"] = i.SourceID
}

func (i instTaskDetail) clone(db *gorm.DB) instTaskDetail {
	i.instTaskDetailDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i instTaskDetail) replaceDB(db *gorm.DB) instTaskDetail {
	i.instTaskDetailDo.ReplaceDB(db)
	return i
}

type instTaskDetailDo struct{ gen.DO }

func (i instTaskDetailDo) Debug() *instTaskDetailDo {
	return i.withDO(i.DO.Debug())
}

func (i instTaskDetailDo) WithContext(ctx context.Context) *instTaskDetailDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i instTaskDetailDo) ReadDB() *instTaskDetailDo {
	return i.Clauses(dbresolver.Read)
}

func (i instTaskDetailDo) WriteDB() *instTaskDetailDo {
	return i.Clauses(dbresolver.Write)
}

func (i instTaskDetailDo) Session(config *gorm.Session) *instTaskDetailDo {
	return i.withDO(i.DO.Session(config))
}

func (i instTaskDetailDo) Clauses(conds ...clause.Expression) *instTaskDetailDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i instTaskDetailDo) Returning(value interface{}, columns ...string) *instTaskDetailDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i instTaskDetailDo) Not(conds ...gen.Condition) *instTaskDetailDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i instTaskDetailDo) Or(conds ...gen.Condition) *instTaskDetailDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i instTaskDetailDo) Select(conds ...field.Expr) *instTaskDetailDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i instTaskDetailDo) Where(conds ...gen.Condition) *instTaskDetailDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i instTaskDetailDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *instTaskDetailDo {
	return i.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (i instTaskDetailDo) Order(conds ...field.Expr) *instTaskDetailDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i instTaskDetailDo) Distinct(cols ...field.Expr) *instTaskDetailDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i instTaskDetailDo) Omit(cols ...field.Expr) *instTaskDetailDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i instTaskDetailDo) Join(table schema.Tabler, on ...field.Expr) *instTaskDetailDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i instTaskDetailDo) LeftJoin(table schema.Tabler, on ...field.Expr) *instTaskDetailDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i instTaskDetailDo) RightJoin(table schema.Tabler, on ...field.Expr) *instTaskDetailDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i instTaskDetailDo) Group(cols ...field.Expr) *instTaskDetailDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i instTaskDetailDo) Having(conds ...gen.Condition) *instTaskDetailDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i instTaskDetailDo) Limit(limit int) *instTaskDetailDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i instTaskDetailDo) Offset(offset int) *instTaskDetailDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i instTaskDetailDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *instTaskDetailDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i instTaskDetailDo) Unscoped() *instTaskDetailDo {
	return i.withDO(i.DO.Unscoped())
}

func (i instTaskDetailDo) Create(values ...*model.InstTaskDetail) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i instTaskDetailDo) CreateInBatches(values []*model.InstTaskDetail, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i instTaskDetailDo) Save(values ...*model.InstTaskDetail) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i instTaskDetailDo) First() (*model.InstTaskDetail, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstTaskDetail), nil
	}
}

func (i instTaskDetailDo) Take() (*model.InstTaskDetail, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstTaskDetail), nil
	}
}

func (i instTaskDetailDo) Last() (*model.InstTaskDetail, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstTaskDetail), nil
	}
}

func (i instTaskDetailDo) Find() ([]*model.InstTaskDetail, error) {
	result, err := i.DO.Find()
	return result.([]*model.InstTaskDetail), err
}

func (i instTaskDetailDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.InstTaskDetail, err error) {
	buf := make([]*model.InstTaskDetail, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i instTaskDetailDo) FindInBatches(result *[]*model.InstTaskDetail, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i instTaskDetailDo) Attrs(attrs ...field.AssignExpr) *instTaskDetailDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i instTaskDetailDo) Assign(attrs ...field.AssignExpr) *instTaskDetailDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i instTaskDetailDo) Joins(fields ...field.RelationField) *instTaskDetailDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i instTaskDetailDo) Preload(fields ...field.RelationField) *instTaskDetailDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i instTaskDetailDo) FirstOrInit() (*model.InstTaskDetail, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstTaskDetail), nil
	}
}

func (i instTaskDetailDo) FirstOrCreate() (*model.InstTaskDetail, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstTaskDetail), nil
	}
}

func (i instTaskDetailDo) FindByPage(offset int, limit int) (result []*model.InstTaskDetail, count int64, err error) {
	result, err = i.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = i.Offset(-1).Limit(-1).Count()
	return
}

func (i instTaskDetailDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i instTaskDetailDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i instTaskDetailDo) Delete(models ...*model.InstTaskDetail) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *instTaskDetailDo) withDO(do gen.Dao) *instTaskDetailDo {
	i.DO = *do.(*gen.DO)
	return i
}