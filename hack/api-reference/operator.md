<p>Packages:</p>
<ul>
<li>
<a href="#agent-sandbox.extensions.gardener.cloud%2fv1alpha1">agent-sandbox.extensions.gardener.cloud/v1alpha1</a>
</li>
</ul>
<h2 id="agent-sandbox.extensions.gardener.cloud/v1alpha1">agent-sandbox.extensions.gardener.cloud/v1alpha1</h2>
<p>
<p>Package v1alpha1 contains the shoot agent-sandbox extension configuration.</p>
</p>
Resource Types:
<ul><li>
<a href="#agent-sandbox.extensions.gardener.cloud/v1alpha1.AgentSandbox">AgentSandbox</a>
</li></ul>
<h3 id="agent-sandbox.extensions.gardener.cloud/v1alpha1.AgentSandbox">AgentSandbox
</h3>
<p>
<p>AgentSandbox contains the configuration for the agent-sandbox controller.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code></br>
string</td>
<td>
<code>
agent-sandbox.extensions.gardener.cloud/v1alpha1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code></br>
string
</td>
<td><code>AgentSandbox</code></td>
</tr>
<tr>
<td>
<code>enableExtensions</code></br>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>EnableExtensions enables extension CRDs (SandboxClaim, SandboxTemplate, SandboxWarmPool) and their RBAC</p>
</td>
</tr>
</tbody>
</table>
<hr/>
<p><em>
Generated with <a href="https://github.com/ahmetb/gen-crd-api-reference-docs">gen-crd-api-reference-docs</a>
</em></p>
