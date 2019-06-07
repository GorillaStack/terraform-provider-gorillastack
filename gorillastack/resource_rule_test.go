package gorillastack

// import (
// 	"fmt"
// 	"testing"

// 	"os"
// 	"regexp"

// 	"github.com/hashicorp/terraform/helper/acctest"
// 	"github.com/hashicorp/terraform/helper/resource"
// 	"github.com/hashicorp/terraform/terraform"
// )

// var certificateArnRegex = regexp.MustCompile(`^arn:aws:acm:[^:]+:[^:]+:certificate/.+$`)

// const REQUIRED_ENV_VARS = []string{
// 	"GORILLASTACK_ROOT_URL",
// 	"GORILLASTACK_API_KEY",
// }

// func getEnv(val string) string {
// 	if v := os.Getenv(val); v == "" {
// 		t.Fatal("%s is not available", val)
// 	}
// 	return os.Getenv(v)
// }
// func checkEnv(t *testing.T) {
// 	for _, required := range REQUIRED_ENV_VARS {
// 		if v := os.Getenv(required); v == "" {
// 			t.Fatal("%s is required in the environment", required)
// 		}
// 	}
// }

// func TestAccRule3partaws(t *testing.T) {
// 	rootDomain := testAccAwsAcmCertificateDomainFromEnv(t)

// 	rInt1 := acctest.RandInt()

// 	domain := fmt.Sprintf("tf-acc-%d.%s", rInt1, rootDomain)

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck:     func() { testAccPreCheck(t) },
// 		Providers:    testAccProviders,
// 		CheckDestroy: testAccCheckAcmCertificateDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: (domain, acm.ValidationMethodEmail),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestMatchResourceAttr("aws_acm_certificate.cert", "arn", certificateArnRegex),
// 					resource.TestCheckResourceAttr("aws_acm_certificate.cert", "domain_name", domain),
// 					resource.TestCheckResourceAttr("aws_acm_certificate.cert", "domain_validation_options.#", "0"),
// 					resource.TestCheckResourceAttr("aws_acm_certificate.cert", "subject_alternative_names.#", "0"),
// 					resource.TestMatchResourceAttr("aws_acm_certificate.cert", "validation_emails.0", regexp.MustCompile(`^[^@]+@.+$`)),
// 					resource.TestCheckResourceAttr("aws_acm_certificate.cert", "validation_method", acm.ValidationMethodEmail),
// 				),
// 			},
// 			{
// 				ResourceName:      "aws_acm_certificate.cert",
// 				ImportState:       true,
// 				ImportStateVerify: true,
// 			},
// 		},
// 	})
// }

// func testAccCheckAcmCertificateDestroy(s *terraform.State) error {
// 	acmconn := testAccProvider.Meta().(*client).acmconn

// 	for _, rs := range s.RootModule().Resources {
// 		if rs.Type != "aws_acm_certificate" {
// 			continue
// 		}
// 		_, err := acmconn.DescribeCertificate(&acm.DescribeCertificateInput{
// 			CertificateArn: aws.String(rs.Primary.ID),
// 		})

// 		if err == nil {
// 			return fmt.Errorf("Rule still exists.")
// 		}

// 		// Verify the error is what we want
// 		if !isAWSErr(err, acm.ErrCodeResourceNotFoundException, "") {
// 			return err
// 		}
// 	}

// 	return nil
// }
