package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgressConn struct {
	connectUri string
	config     *pgxpool.Config
	pool       *pgxpool.Pool
}

func New() *PostgressConn {
	return &PostgressConn{}
}

func (pr *PostgressConn) GetConn(ctx context.Context) (*pgxpool.Conn, error) {

	ctxReq, done := context.WithTimeout(ctx, time.Second*(30))
	defer done()
	conn, err := pr.pool.Acquire(ctxReq)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (pr *PostgressConn) Connect(ctx context.Context, connString string) (func(), error) {
	var runtimeParams map[string]string = make(map[string]string)
	runtimeParams["application_name"] = "Servidor_Test"

	if pr.config == nil {
		config, err := pgxpool.ParseConfig(connString)
		if err != nil { // Connection string Parse error, exit
			return nil, err
		}
		pr.config = config
		pr.connectUri = connString
	}

	ctxReq, done := context.WithTimeout(ctx, time.Second*(30))
	defer done()

	pool, err := pgxpool.ConnectConfig(ctxReq, pr.config)
	if err != nil {
		return nil, err
	}
	pr.pool = pool // Store Connection in the Repository

	return func() { // Return a closer
		pool.Close()
	}, nil
}
