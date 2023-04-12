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

func newProcessDefNodeHandler(db *gorm.DB, opts ...gen.DOOption) processDefNodeHandler {
	_processDefNodeHandler := processDefNodeHandler{}

	_processDefNodeHandler.processDefNodeHandlerDo.UseDB(db, opts...)
	_processDefNodeHandler.processDefNodeHandlerDo.UseModel(&model.ProcessDefNodeHandler{})

	tableName := _processDefNodeHandler.processDefNodeHandlerDo.TableName()
	_processDefNodeHandler.ALL = field.NewAsterisk(tableName)
	_processDefNodeHandler.ID = field.NewInt64(tableName, "id")
	_processDefNodeHandler.ProcessDefID = field.NewString(tableName, "process_def_id")
	_processDefNodeHandler.NodeID = field.NewString(tableName, "node_id")
	_processDefNodeHandler.HandlerName = field.NewString(tableName, "handler_name")
	_processDefNodeHandler.HandlerType = field.NewInt32(tableName, "handler_type")
	_processDefNodeHandler.HandlerID = field.NewString(tableName, "handler_id")
	_processDefNodeHandler.HandlerSort = field.NewInt32(tableName, "handler_sort")
	_processDefNodeHandler.ObjData = field.NewString(tableName, "obj_data")
	_processDefNodeHandler.CreateTime = field.NewTime(tableName, "create_time")
	_processDefNodeHandler.CreateUser = field.NewString(tableName, "create_user")
	_processDefNodeHandler.UpdateTime = field.NewTime(tableName, "update_time")
	_processDefNodeHandler.UpdateUser = field.NewString(tableName, "update_user")

	_processDefNodeHandler.fillFieldMap()

	return _processDefNodeHandler
}

type processDefNodeHandler struct {
	processDefNodeHandlerDo processDefNodeHandlerDo

	ALL          field.Asterisk
	ID           field.Int64  // 唯一id
	ProcessDefID field.String // 流程定义id
	NodeID       field.String // 节点id
	HandlerName  field.String // 处理人名称
	HandlerType  field.Int32  // 处理人类型【1：用户；2：部门；3：相对岗位；4：表单控件；5：部门岗位】
	HandlerID    field.String // 处理人对象id;处理对象的id，根据处理人类型区分，如果操作员id、部门id等
	HandlerSort  field.Int32  // 处理人顺序;正序排序
	ObjData      field.String // 对象数据;依据处理人类型取值，相对岗位和表单控件使用该字段存json数据
	CreateTime   field.Time   // 创建时间
	CreateUser   field.String // 创建人
	UpdateTime   field.Time   // 更新时间
	UpdateUser   field.String // 更新人

	fieldMap map[string]field.Expr
}

func (p processDefNodeHandler) Table(newTableName string) *processDefNodeHandler {
	p.processDefNodeHandlerDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p processDefNodeHandler) As(alias string) *processDefNodeHandler {
	p.processDefNodeHandlerDo.DO = *(p.processDefNodeHandlerDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *processDefNodeHandler) updateTableName(table string) *processDefNodeHandler {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.ProcessDefID = field.NewString(table, "process_def_id")
	p.NodeID = field.NewString(table, "node_id")
	p.HandlerName = field.NewString(table, "handler_name")
	p.HandlerType = field.NewInt32(table, "handler_type")
	p.HandlerID = field.NewString(table, "handler_id")
	p.HandlerSort = field.NewInt32(table, "handler_sort")
	p.ObjData = field.NewString(table, "obj_data")
	p.CreateTime = field.NewTime(table, "create_time")
	p.CreateUser = field.NewString(table, "create_user")
	p.UpdateTime = field.NewTime(table, "update_time")
	p.UpdateUser = field.NewString(table, "update_user")

	p.fillFieldMap()

	return p
}

func (p *processDefNodeHandler) WithContext(ctx context.Context) *processDefNodeHandlerDo {
	return p.processDefNodeHandlerDo.WithContext(ctx)
}

func (p processDefNodeHandler) TableName() string { return p.processDefNodeHandlerDo.TableName() }

func (p processDefNodeHandler) Alias() string { return p.processDefNodeHandlerDo.Alias() }

func (p *processDefNodeHandler) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *processDefNodeHandler) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 12)
	p.fieldMap["id"] = p.ID
	p.fieldMap["process_def_id"] = p.ProcessDefID
	p.fieldMap["node_id"] = p.NodeID
	p.fieldMap["handler_name"] = p.HandlerName
	p.fieldMap["handler_type"] = p.HandlerType
	p.fieldMap["handler_id"] = p.HandlerID
	p.fieldMap["handler_sort"] = p.HandlerSort
	p.fieldMap["obj_data"] = p.ObjData
	p.fieldMap["create_time"] = p.CreateTime
	p.fieldMap["create_user"] = p.CreateUser
	p.fieldMap["update_time"] = p.UpdateTime
	p.fieldMap["update_user"] = p.UpdateUser
}

func (p processDefNodeHandler) clone(db *gorm.DB) processDefNodeHandler {
	p.processDefNodeHandlerDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p processDefNodeHandler) replaceDB(db *gorm.DB) processDefNodeHandler {
	p.processDefNodeHandlerDo.ReplaceDB(db)
	return p
}

type processDefNodeHandlerDo struct{ gen.DO }

func (p processDefNodeHandlerDo) Debug() *processDefNodeHandlerDo {
	return p.withDO(p.DO.Debug())
}

func (p processDefNodeHandlerDo) WithContext(ctx context.Context) *processDefNodeHandlerDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p processDefNodeHandlerDo) ReadDB() *processDefNodeHandlerDo {
	return p.Clauses(dbresolver.Read)
}

func (p processDefNodeHandlerDo) WriteDB() *processDefNodeHandlerDo {
	return p.Clauses(dbresolver.Write)
}

func (p processDefNodeHandlerDo) Session(config *gorm.Session) *processDefNodeHandlerDo {
	return p.withDO(p.DO.Session(config))
}

func (p processDefNodeHandlerDo) Clauses(conds ...clause.Expression) *processDefNodeHandlerDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p processDefNodeHandlerDo) Returning(value interface{}, columns ...string) *processDefNodeHandlerDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p processDefNodeHandlerDo) Not(conds ...gen.Condition) *processDefNodeHandlerDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p processDefNodeHandlerDo) Or(conds ...gen.Condition) *processDefNodeHandlerDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p processDefNodeHandlerDo) Select(conds ...field.Expr) *processDefNodeHandlerDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p processDefNodeHandlerDo) Where(conds ...gen.Condition) *processDefNodeHandlerDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p processDefNodeHandlerDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *processDefNodeHandlerDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p processDefNodeHandlerDo) Order(conds ...field.Expr) *processDefNodeHandlerDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p processDefNodeHandlerDo) Distinct(cols ...field.Expr) *processDefNodeHandlerDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p processDefNodeHandlerDo) Omit(cols ...field.Expr) *processDefNodeHandlerDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p processDefNodeHandlerDo) Join(table schema.Tabler, on ...field.Expr) *processDefNodeHandlerDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p processDefNodeHandlerDo) LeftJoin(table schema.Tabler, on ...field.Expr) *processDefNodeHandlerDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p processDefNodeHandlerDo) RightJoin(table schema.Tabler, on ...field.Expr) *processDefNodeHandlerDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p processDefNodeHandlerDo) Group(cols ...field.Expr) *processDefNodeHandlerDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p processDefNodeHandlerDo) Having(conds ...gen.Condition) *processDefNodeHandlerDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p processDefNodeHandlerDo) Limit(limit int) *processDefNodeHandlerDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p processDefNodeHandlerDo) Offset(offset int) *processDefNodeHandlerDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p processDefNodeHandlerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *processDefNodeHandlerDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p processDefNodeHandlerDo) Unscoped() *processDefNodeHandlerDo {
	return p.withDO(p.DO.Unscoped())
}

func (p processDefNodeHandlerDo) Create(values ...*model.ProcessDefNodeHandler) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p processDefNodeHandlerDo) CreateInBatches(values []*model.ProcessDefNodeHandler, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p processDefNodeHandlerDo) Save(values ...*model.ProcessDefNodeHandler) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p processDefNodeHandlerDo) First() (*model.ProcessDefNodeHandler, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProcessDefNodeHandler), nil
	}
}

func (p processDefNodeHandlerDo) Take() (*model.ProcessDefNodeHandler, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProcessDefNodeHandler), nil
	}
}

func (p processDefNodeHandlerDo) Last() (*model.ProcessDefNodeHandler, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProcessDefNodeHandler), nil
	}
}

func (p processDefNodeHandlerDo) Find() ([]*model.ProcessDefNodeHandler, error) {
	result, err := p.DO.Find()
	return result.([]*model.ProcessDefNodeHandler), err
}

func (p processDefNodeHandlerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProcessDefNodeHandler, err error) {
	buf := make([]*model.ProcessDefNodeHandler, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p processDefNodeHandlerDo) FindInBatches(result *[]*model.ProcessDefNodeHandler, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p processDefNodeHandlerDo) Attrs(attrs ...field.AssignExpr) *processDefNodeHandlerDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p processDefNodeHandlerDo) Assign(attrs ...field.AssignExpr) *processDefNodeHandlerDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p processDefNodeHandlerDo) Joins(fields ...field.RelationField) *processDefNodeHandlerDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p processDefNodeHandlerDo) Preload(fields ...field.RelationField) *processDefNodeHandlerDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p processDefNodeHandlerDo) FirstOrInit() (*model.ProcessDefNodeHandler, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProcessDefNodeHandler), nil
	}
}

func (p processDefNodeHandlerDo) FirstOrCreate() (*model.ProcessDefNodeHandler, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProcessDefNodeHandler), nil
	}
}

func (p processDefNodeHandlerDo) FindByPage(offset int, limit int) (result []*model.ProcessDefNodeHandler, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p processDefNodeHandlerDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p processDefNodeHandlerDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p processDefNodeHandlerDo) Delete(models ...*model.ProcessDefNodeHandler) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *processDefNodeHandlerDo) withDO(do gen.Dao) *processDefNodeHandlerDo {
	p.DO = *do.(*gen.DO)
	return p
}
