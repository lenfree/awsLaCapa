package test

import (
        _ "github.com/lenfree/awsLaCapa/routers"
        "net/http"
        "net/http/httptest"
        "path/filepath"
        "runtime"
        "testing"

        "github.com/astaxie/beego"
        . "github.com/smartystreets/goconvey/convey"
)

func init() {
        _, file, _, _ := runtime.Caller(1)
        apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
        beego.TestBeegoInit(apppath)
}

func TestGetIAMUsers(t *testing.T) {
        r, _ := http.NewRequest("GET", "/v1/iam/users", nil)
        w := httptest.NewRecorder()
        beego.BeeApp.Handlers.ServeHTTP(w, r)

        beego.Trace("testing", "TestGetIAMUsers", "Code[%d]\n%s", w.Code, w.Body.String())

        Convey("Subject: Test GetIAMUsers Endpoint\n", t, func() {
                Convey("Status Code Should Be 200", func() {
                        So(w.Code, ShouldEqual, 200)
                })
        })
}

func TestGetIAMGroups(t *testing.T) {
        r, _ := http.NewRequest("GET", "/v1/iam/groups", nil)
        w := httptest.NewRecorder()
        beego.BeeApp.Handlers.ServeHTTP(w, r)

        beego.Trace("testing", "TestGetIAMGroups", "Code[%d]\n%s", w.Code, w.Body.String())

        Convey("Subject: Test GetIAMGroups Endpoint\n", t, func() {
                Convey("Status Code Should Be 200", func() {
                        So(w.Code, ShouldEqual, 200)
                })
        })
}

func TestGetIAMUserGroups(t *testing.T) {
        r, _ := http.NewRequest("GET", "/v1/iam/user_groups", nil)
        w := httptest.NewRecorder()
        beego.BeeApp.Handlers.ServeHTTP(w, r)

        beego.Trace("testing", "TestGetIAMUserGroups", "Code[%d]\n%s", w.Code, w.Body.String())

        Convey("Subject: Test GetIAMUserGroups Endpoint\n", t, func() {
                Convey("Status Code Should Be 200", func() {
                        So(w.Code, ShouldEqual, 200)
                })
        })
}

func TestGetS3(t *testing.T) {
        r, _ := http.NewRequest("GET", "/v1/s3", nil)
        w := httptest.NewRecorder()
        beego.BeeApp.Handlers.ServeHTTP(w, r)

        beego.Trace("testing", "TestGetS3", "Code[%d]\n%s", w.Code, w.Body.String())

        Convey("Subject: Test GetS3 Endpoint\n", t, func() {
                Convey("Status Code Should Be 200", func() {
                        So(w.Code, ShouldEqual, 200)
                })
        })
}

func TestGetVPC(t *testing.T) {
        r, _ := http.NewRequest("GET", "/v1/vpc", nil)
        w := httptest.NewRecorder()
        beego.BeeApp.Handlers.ServeHTTP(w, r)

        beego.Trace("testing", "TestGetVPC", "Code[%d]\n%s", w.Code, w.Body.String())

        Convey("Subject: Test VPCs Endpoint\n", t, func() {
                Convey("Status Code Should Be 200", func() {
                        So(w.Code, ShouldEqual, 200)
                })
        })
}

func TestGetVPCPeeringConnections(t *testing.T) {
        r, _ := http.NewRequest("GET", "/v1/vpc/vpc_peering_connections", nil)
        w := httptest.NewRecorder()
        beego.BeeApp.Handlers.ServeHTTP(w, r)

        beego.Trace("testing", "TestGetVPCPeeringConnections", "Code[%d]\n%s", w.Code, w.Body.String())

        Convey("Subject: Test VPC Peering Connections Endpoint\n", t, func() {
                Convey("Status Code Should Be 200", func() {
                        So(w.Code, ShouldEqual, 200)
                })
        })
}

func TestS3ListObjects(t *testing.T) {
        r, _ := http.NewRequest("GET", "/v1/s3/mybucket", nil)
        w := httptest.NewRecorder()
        beego.BeeApp.Handlers.ServeHTTP(w, r)

        beego.Trace("testing", "TestS3ListObjects", "Code[%d]\n%s", w.Code, w.Body.String())

        Convey("Subject: Test S3 Bucket Endpoint\n", t, func() {
                Convey("Status Code Should Be 200", func() {
                        So(w.Code, ShouldEqual, 200)
                })
        })
}

func TestS3GetObjectByKey(t *testing.T) {
        r, _ := http.NewRequest("GET", "/v1/s3/mybucket/directory/key", nil)
        w := httptest.NewRecorder()
        beego.BeeApp.Handlers.ServeHTTP(w, r)

        beego.Trace("testing", "TestS3GetObjectByKey", "Code[%d]\n%s", w.Code, w.Body.String())

        Convey("Subject: Test S3 Bucket Object Key Endpoint\n", t, func() {
                Convey("Status Code Should Be 200", func() {
                        So(w.Code, ShouldEqual, 200)
                })
        })

        r, _ = http.NewRequest("GET", "/v1/s3/mybucket/key", nil)
        w = httptest.NewRecorder()
        beego.BeeApp.Handlers.ServeHTTP(w, r)

        beego.Trace("testing", "TestS3GetObjectByKeyNotFound", "Code[%d]\n%s", w.Code)

        Convey("Subject: Test S3 Bucket Object Key Endpoint With Incorrect Directory Path Or Prefix\n", t, func() {
                Convey("Status Code Should Be 404", func() {
                        So(w.Code, ShouldEqual, 404)
                })
        })
}

func TestGetRoute53(t *testing.T) {
        r, _ := http.NewRequest("GET", "/v1/route53", nil)
        w := httptest.NewRecorder()
        beego.BeeApp.Handlers.ServeHTTP(w, r)

        beego.Trace("testing", "TestGetRoute53", "Code[%d]\n%s", w.Code, w.Body.String())

        Convey("Subject: Test Route53 Endpoint\n", t, func() {
                Convey("Status Code Should Be 200", func() {
                        So(w.Code, ShouldEqual, 200)
                })
        })
}

func TestGetRoute53RRSet(t *testing.T) {
        r, _ := http.NewRequest("GET", "/v1/route53/hostedzone/id-01", nil)
        w := httptest.NewRecorder()
        beego.BeeApp.Handlers.ServeHTTP(w, r)

        beego.Trace("testing", "TestGetRoute53RRSet", "Code[%d]\n%s", w.Code, w.Body.String())

        Convey("Subject: Test Route53 RR Set Endpoint\n", t, func() {
                Convey("Status Code Should Be 200", func() {
                        So(w.Code, ShouldEqual, 200)
                })
        })
}

func TestGetEC2Instances(t *testing.T) {
        r, _ := http.NewRequest("GET", "/v1/ec2", nil)
        w := httptest.NewRecorder()
        beego.BeeApp.Handlers.ServeHTTP(w, r)

        beego.Trace("testing", "TestGetEC2Instances", "Code[%d]\n%s", w.Code, w.Body.String())

        Convey("Subject: Test EC2 Instace Endpoint\n", t, func() {
                Convey("Status Code Should Be 200", func() {
                        So(w.Code, ShouldEqual, 200)
                })
        })
}

func TestGetDHCPOptionSet(t *testing.T) {
        r, _ := http.NewRequest("GET", "/v1/dhcp", nil)
        w := httptest.NewRecorder()
        beego.BeeApp.Handlers.ServeHTTP(w, r)

        beego.Trace("testing", "TestGetDHCPOptionSet", "Code[%d]\n%s", w.Code, w.Body.String())

        Convey("Subject: Test DHCP Endpoint\n", t, func() {
                Convey("Status Code Should Be 200", func() {
                        So(w.Code, ShouldEqual, 200)
                })
        })
}

func TestGetSecurityGroups(t *testing.T) {
        r, _ := http.NewRequest("GET", "/v1/security_groups", nil)
        w := httptest.NewRecorder()
        beego.BeeApp.Handlers.ServeHTTP(w, r)

        beego.Trace("testing", "TestGetSecurityGroups", "Code[%d]\n%s", w.Code, w.Body.String())

        Convey("Subject: Test Security Groups Endpoint\n", t, func() {
                Convey("Status Code Should Be 200", func() {
                        So(w.Code, ShouldEqual, 200)
                })
        })
}
