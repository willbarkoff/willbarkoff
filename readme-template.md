#### 👋 {{.Greeting}}
##### I'm Will Barkoff, a student by day and a developer by night. I'm indecisive at dusk. 

---

Here are some things I've written recently.
{{range $index, $post := .Posts}}
- [{{print $post.Title}}]({{print $post.URL}}) &mdash; _{{print $post.Date}}_
{{end}}

Read more at my website, [willbarkoff.dev](https://willbarkoff.dev).

---
_This page was last generated on {{.GeneratedAt}}_