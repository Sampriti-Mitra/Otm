package constants

const (
	// Input: key where input data is stored as map
	Input = "input"

	// RequestID: holds the unique request identifier for the request
	RequestID = "request_id"

	// ContainerID: k8s container id
	ContainerID = "container_id"

	// CommitID - git commit hash
	CommitID = "commit_id"

	// GitCommitHash - git commit hash env key
	GitCommitHash = "GIT_COMMIT_HASH"

	// Hostname: k8s container id env key
	Hostname = "HOSTNAME"

	// DB: key to hold the database instance
	DB = "database"

	// Logger: holds entry in context
	LOGGER = "logger"

	// ExternalRequestID: holds the unique external request identifier
	ExternalRequestID = "external_request_id"

	// TaskID: is the key to hold the task id of the process
	TaskID = "task_id"

	// AppMode: used it identify the application to run on debug mode
	AppMode = "APP_MODE"

	// Response: keys where response has to be written
	Request = "request"

	// Mode : cli flag to specify env for migrations
	Mode = "mode"

	//Service: map which contains all the system/pod information
	Service = "service"

	//Context: map contains tracedata
	Context = "context"

	// TaskID: is the key to hold the task id of the process
	X_RAZORPAY_TASK_ID = "X-Razorpay-TaskId"

	TestData = "TestData"

	MaxRecursionAllowed = 6

	RazorXClient = "RAZORX"

	RazorXVariantOn = "on"
)
const (
	Config = "config"
)
