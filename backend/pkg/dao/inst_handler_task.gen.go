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

	"github.com/wegoteam/weflow/pkg/model"
)

func newInstHandlerTask(db *gorm.DB, opts ...gen.DOOption) instHandlerTask {
	_instHandlerTask := instHandlerTask{}

	_instHandlerTask.instHandlerTaskDo.UseDB(db, opts...)
	_instHandlerTask.instHandlerTaskDo.UseModel(&model.InstHandlerTask{})

	tableName := _instHandlerTask.instHandlerTaskDo.TableName()
	_instHandlerTask.ALL = field.NewAsterisk(tableName)
	_instHandlerTask.ID = field.NewInt64(tableName, "id")
	_instHandlerTask.InstTaskID = field.NewString(tableName, "inst_task_id")
	_instHandlerTask.NodeTaskID = field.NewString(tableName, "node_task_id")
	_instHandlerTask.NodeID = field.NewString(tableName, "node_id")
	_instHandlerTask.HandlerTaskID = field.NewString(tableName, "handler_task_id")
	_instHandlerTask.NodeHandlerID = field.NewString(tableName, "node_handler_id")
	_instHandlerTask.NodeHandlerName = field.NewString(tableName, "node_handler_name")
	_instHandlerTask.NodeHandlerType = field.NewInt32(tableName, "node_handler_type")
	_instHandlerTask.OpOrigin = field.NewInt32(tableName, "op_origin")
	_instHandlerTask.TimeLimit = field.NewInt64(tableName, "time_limit")
	_instHandlerTask.Status = field.NewInt32(tableName, "status")
	_instHandlerTask.CreateTime = field.NewTime(tableName, "create_time")
	_instHandlerTask.UpdateTime = field.NewTime(tableName, "update_time")
	_instHandlerTask.HandleTime = field.NewTime(tableName, "handle_time")
	_instHandlerTask.OpUserID = field.NewString(tableName, "op_user_id")
	_instHandlerTask.OpUserName = field.NewString(tableName, "op_user_name")
	_instHandlerTask.HandlerSort = field.NewInt32(tableName, "handler_sort")
	_instHandlerTask.Opinion = field.NewInt32(tableName, "opinion")
	_instHandlerTask.OpinionDesc = field.NewString(tableName, "opinion_desc")

	_instHandlerTask.fillFieldMap()

	return _instHandlerTask
}

type instHandlerTask struct {
	instHandlerTaskDo instHandlerTaskDo

	ALL             field.Asterisk
	ID              field.Int64  // 唯一id
	InstTaskID      field.String // 实例任务id
	NodeTaskID      field.String // 节点任务id
	NodeID          field.String // 节点任务id
	HandlerTaskID   field.String // 处理人任务id
	NodeHandlerID   field.String // 节点处理人id
	NodeHandlerName field.String // 处理人名称
	NodeHandlerType field.Int32  // 处理人类型【1：操作员；2：部门；3：相对岗位；4：表单控件；5：角色；6：岗位；7：组织；8：自定义】
	OpOrigin        field.Int32  // 操作来源【1：正常；2：加签】
	TimeLimit       field.Int64  // 处理期限;格式：yyyymmddhhmm 可直接指定到期限的具体时间，期限支持到分钟； 0表示无期限
	Status          field.Int32  // 任务状态【1：处理中；2：完成；3：回退；4：终止】
	CreateTime      field.Time   // 创建时间
	UpdateTime      field.Time   // 更新时间
	HandleTime      field.Time   // 处理时间
	OpUserID        field.String // 操作用户id
	OpUserName      field.String // 操作用户名称
	HandlerSort     field.Int32  // 处理人排序;处理人当前的处理排序
	Opinion         field.Int32  // 处理意见【1：未发表；2：已阅；3：同意；4：不同意】
	OpinionDesc     field.String // 处理意见描述

	fieldMap map[string]field.Expr
}

func (i instHandlerTask) Table(newTableName string) *instHandlerTask {
	i.instHandlerTaskDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i instHandlerTask) As(alias string) *instHandlerTask {
	i.instHandlerTaskDo.DO = *(i.instHandlerTaskDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *instHandlerTask) updateTableName(table string) *instHandlerTask {
	i.ALL = field.NewAsterisk(table)
	i.ID = field.NewInt64(table, "id")
	i.InstTaskID = field.NewString(table, "inst_task_id")
	i.NodeTaskID = field.NewString(table, "node_task_id")
	i.NodeID = field.NewString(table, "node_id")
	i.HandlerTaskID = field.NewString(table, "handler_task_id")
	i.NodeHandlerID = field.NewString(table, "node_handler_id")
	i.NodeHandlerName = field.NewString(table, "node_handler_name")
	i.NodeHandlerType = field.NewInt32(table, "node_handler_type")
	i.OpOrigin = field.NewInt32(table, "op_origin")
	i.TimeLimit = field.NewInt64(table, "time_limit")
	i.Status = field.NewInt32(table, "status")
	i.CreateTime = field.NewTime(table, "create_time")
	i.UpdateTime = field.NewTime(table, "update_time")
	i.HandleTime = field.NewTime(table, "handle_time")
	i.OpUserID = field.NewString(table, "op_user_id")
	i.OpUserName = field.NewString(table, "op_user_name")
	i.HandlerSort = field.NewInt32(table, "handler_sort")
	i.Opinion = field.NewInt32(table, "opinion")
	i.OpinionDesc = field.NewString(table, "opinion_desc")

	i.fillFieldMap()

	return i
}

func (i *instHandlerTask) WithContext(ctx context.Context) *instHandlerTaskDo {
	return i.instHandlerTaskDo.WithContext(ctx)
}

func (i instHandlerTask) TableName() string { return i.instHandlerTaskDo.TableName() }

func (i instHandlerTask) Alias() string { return i.instHandlerTaskDo.Alias() }

func (i *instHandlerTask) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *instHandlerTask) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 19)
	i.fieldMap["id"] = i.ID
	i.fieldMap["inst_task_id"] = i.InstTaskID
	i.fieldMap["node_task_id"] = i.NodeTaskID
	i.fieldMap["node_id"] = i.NodeID
	i.fieldMap["handler_task_id"] = i.HandlerTaskID
	i.fieldMap["node_handler_id"] = i.NodeHandlerID
	i.fieldMap["node_handler_name"] = i.NodeHandlerName
	i.fieldMap["node_handler_type"] = i.NodeHandlerType
	i.fieldMap["op_origin"] = i.OpOrigin
	i.fieldMap["time_limit"] = i.TimeLimit
	i.fieldMap["status"] = i.Status
	i.fieldMap["create_time"] = i.CreateTime
	i.fieldMap["update_time"] = i.UpdateTime
	i.fieldMap["handle_time"] = i.HandleTime
	i.fieldMap["op_user_id"] = i.OpUserID
	i.fieldMap["op_user_name"] = i.OpUserName
	i.fieldMap["handler_sort"] = i.HandlerSort
	i.fieldMap["opinion"] = i.Opinion
	i.fieldMap["opinion_desc"] = i.OpinionDesc
}

func (i instHandlerTask) clone(db *gorm.DB) instHandlerTask {
	i.instHandlerTaskDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i instHandlerTask) replaceDB(db *gorm.DB) instHandlerTask {
	i.instHandlerTaskDo.ReplaceDB(db)
	return i
}

type instHandlerTaskDo struct{ gen.DO }

func (i instHandlerTaskDo) Debug() *instHandlerTaskDo {
	return i.withDO(i.DO.Debug())
}

func (i instHandlerTaskDo) WithContext(ctx context.Context) *instHandlerTaskDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i instHandlerTaskDo) ReadDB() *instHandlerTaskDo {
	return i.Clauses(dbresolver.Read)
}

func (i instHandlerTaskDo) WriteDB() *instHandlerTaskDo {
	return i.Clauses(dbresolver.Write)
}

func (i instHandlerTaskDo) Session(config *gorm.Session) *instHandlerTaskDo {
	return i.withDO(i.DO.Session(config))
}

func (i instHandlerTaskDo) Clauses(conds ...clause.Expression) *instHandlerTaskDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i instHandlerTaskDo) Returning(value interface{}, columns ...string) *instHandlerTaskDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i instHandlerTaskDo) Not(conds ...gen.Condition) *instHandlerTaskDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i instHandlerTaskDo) Or(conds ...gen.Condition) *instHandlerTaskDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i instHandlerTaskDo) Select(conds ...field.Expr) *instHandlerTaskDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i instHandlerTaskDo) Where(conds ...gen.Condition) *instHandlerTaskDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i instHandlerTaskDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *instHandlerTaskDo {
	return i.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (i instHandlerTaskDo) Order(conds ...field.Expr) *instHandlerTaskDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i instHandlerTaskDo) Distinct(cols ...field.Expr) *instHandlerTaskDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i instHandlerTaskDo) Omit(cols ...field.Expr) *instHandlerTaskDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i instHandlerTaskDo) Join(table schema.Tabler, on ...field.Expr) *instHandlerTaskDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i instHandlerTaskDo) LeftJoin(table schema.Tabler, on ...field.Expr) *instHandlerTaskDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i instHandlerTaskDo) RightJoin(table schema.Tabler, on ...field.Expr) *instHandlerTaskDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i instHandlerTaskDo) Group(cols ...field.Expr) *instHandlerTaskDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i instHandlerTaskDo) Having(conds ...gen.Condition) *instHandlerTaskDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i instHandlerTaskDo) Limit(limit int) *instHandlerTaskDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i instHandlerTaskDo) Offset(offset int) *instHandlerTaskDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i instHandlerTaskDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *instHandlerTaskDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i instHandlerTaskDo) Unscoped() *instHandlerTaskDo {
	return i.withDO(i.DO.Unscoped())
}

func (i instHandlerTaskDo) Create(values ...*model.InstHandlerTask) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i instHandlerTaskDo) CreateInBatches(values []*model.InstHandlerTask, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i instHandlerTaskDo) Save(values ...*model.InstHandlerTask) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i instHandlerTaskDo) First() (*model.InstHandlerTask, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstHandlerTask), nil
	}
}

func (i instHandlerTaskDo) Take() (*model.InstHandlerTask, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstHandlerTask), nil
	}
}

func (i instHandlerTaskDo) Last() (*model.InstHandlerTask, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstHandlerTask), nil
	}
}

func (i instHandlerTaskDo) Find() ([]*model.InstHandlerTask, error) {
	result, err := i.DO.Find()
	return result.([]*model.InstHandlerTask), err
}

func (i instHandlerTaskDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.InstHandlerTask, err error) {
	buf := make([]*model.InstHandlerTask, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i instHandlerTaskDo) FindInBatches(result *[]*model.InstHandlerTask, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i instHandlerTaskDo) Attrs(attrs ...field.AssignExpr) *instHandlerTaskDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i instHandlerTaskDo) Assign(attrs ...field.AssignExpr) *instHandlerTaskDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i instHandlerTaskDo) Joins(fields ...field.RelationField) *instHandlerTaskDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i instHandlerTaskDo) Preload(fields ...field.RelationField) *instHandlerTaskDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i instHandlerTaskDo) FirstOrInit() (*model.InstHandlerTask, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstHandlerTask), nil
	}
}

func (i instHandlerTaskDo) FirstOrCreate() (*model.InstHandlerTask, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.InstHandlerTask), nil
	}
}

func (i instHandlerTaskDo) FindByPage(offset int, limit int) (result []*model.InstHandlerTask, count int64, err error) {
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

func (i instHandlerTaskDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i instHandlerTaskDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i instHandlerTaskDo) Delete(models ...*model.InstHandlerTask) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *instHandlerTaskDo) withDO(do gen.Dao) *instHandlerTaskDo {
	i.DO = *do.(*gen.DO)
	return i
}
