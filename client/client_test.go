package client

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http/cookiejar"
)

var _ = Describe("#Build", func() {
	Context("#Build", func() {
		It("should return an err if set", func() {
			builder := New()
			builder.err = errors.New("boom!")

			// When
			_, err := builder.Build()

			// Then
			Expect(err).ToNot(BeNil(), "expected to go boom!")
		})
	})

	Context("#WithCookieJarHandler", func() {
		It("should use the custom cookie jar provided", func() {
			jar := MockJarHandler{Id: "1"}

			// When
			client, err := New().WithCookieJarHandler(jar).Build()

			// Then
			Expect(err).To(BeNil(), "expected to construct a client without error")
			Expect(client.Jar).To(Equal(jar), "expected to use our custom CookieJar")
		})
	})

	Context("#WithNoFollowRedirects", func() {
		It("should not follow redirects", func() {
			client, err := New().WithNoFollowRedirects().Build()
			Expect(err).To(BeNil(), "expected no errors")
			Expect(client).ToNot(BeNil(), "expected a client back")

			response, err := client.Get("http://github.com")
			defer response.Body.Close()
			Expect(err).ToNot(BeNil(), "expected an error since this should get redirected")

		})
	})

	Context("#WithCookieJar", func() {
		It("should not create a CookieJar unless #WithCookieJar is specified", func() {
			client, err := New().Build()

			// Then
			Expect(err).To(BeNil(), "expected no error!")
			Expect(client.Jar).To(BeNil(), "expected no CookieJar")
		})

		It("should create client with default cookie jar", func() {
			client, err := New().WithCookieJar(nil).Build()

			// Then
			Expect(err).To(BeNil(), "expected no error!")
			Expect(client).ToNot(BeNil(), "expected client to be returned!")
			Expect(client.Jar).ToNot(BeNil(), "expected client.jar to have been populated!")
		})

		It("should use the options when provided", func() {
			publicSuffixList := PSL{}
			options := cookiejar.Options{PublicSuffixList: publicSuffixList}

			// When
			client, err := New().WithCookieJar(&options).Build()

			// Then
			Expect(err).To(BeNil(), "expected no error!")
			Expect(client).ToNot(BeNil(), "expected client to be returned!")
			Expect(client.Jar).ToNot(BeNil(), "expected client.jar to have been populated!")
			// todo - how to test options has been assigned
		})
	})
})
