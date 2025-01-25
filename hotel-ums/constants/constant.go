package constants

const (
	RedisKeyUserEmail                 = "user:email:%s"
	RedisKeyUserUsername              = "user:username:%s"
	RedisKeyUserID                    = "user:id:%s"
	RedisKeyAllUsers                  = "users"
	RedisKeyUserSessionByToken        = "user:session:%s"
	RedisKeyUserSessionByRefreshToken = "user:session:%s"
)
