// Code generated by qtc from "workload_report.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line pkg/report/templates/workload_report.qtpl:1
package templates

//line pkg/report/templates/workload_report.qtpl:1
import "github.com/aquasecurity/starboard/pkg/apis/aquasecurity/v1alpha1"

//line pkg/report/templates/workload_report.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line pkg/report/templates/workload_report.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line pkg/report/templates/workload_report.qtpl:3
func (p *WorkloadReport) StreamTitle(qw422016 *qt422016.Writer) {
//line pkg/report/templates/workload_report.qtpl:3
	qw422016.N().S(`
Aqua Starboard Workload Security Report - `)
//line pkg/report/templates/workload_report.qtpl:4
	qw422016.E().S(p.Workload.Namespace)
//line pkg/report/templates/workload_report.qtpl:4
	qw422016.N().S(`/`)
//line pkg/report/templates/workload_report.qtpl:4
	qw422016.E().V(p.Workload.Kind)
//line pkg/report/templates/workload_report.qtpl:4
	qw422016.N().S(`/`)
//line pkg/report/templates/workload_report.qtpl:4
	qw422016.E().S(p.Workload.Name)
//line pkg/report/templates/workload_report.qtpl:4
	qw422016.N().S(`
`)
//line pkg/report/templates/workload_report.qtpl:5
}

//line pkg/report/templates/workload_report.qtpl:5
func (p *WorkloadReport) WriteTitle(qq422016 qtio422016.Writer) {
//line pkg/report/templates/workload_report.qtpl:5
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pkg/report/templates/workload_report.qtpl:5
	p.StreamTitle(qw422016)
//line pkg/report/templates/workload_report.qtpl:5
	qt422016.ReleaseWriter(qw422016)
//line pkg/report/templates/workload_report.qtpl:5
}

//line pkg/report/templates/workload_report.qtpl:5
func (p *WorkloadReport) Title() string {
//line pkg/report/templates/workload_report.qtpl:5
	qb422016 := qt422016.AcquireByteBuffer()
//line pkg/report/templates/workload_report.qtpl:5
	p.WriteTitle(qb422016)
//line pkg/report/templates/workload_report.qtpl:5
	qs422016 := string(qb422016.B)
//line pkg/report/templates/workload_report.qtpl:5
	qt422016.ReleaseByteBuffer(qb422016)
//line pkg/report/templates/workload_report.qtpl:5
	return qs422016
//line pkg/report/templates/workload_report.qtpl:5
}

// TODO Check if we need that summary logic. If yes move this logic out of HTML templating.

//line pkg/report/templates/workload_report.qtpl:9
func (p *WorkloadReport) GetMergedVulnsSummary() v1alpha1.VulnerabilitySummary {
	merged := v1alpha1.VulnerabilitySummary{}
	for _, report := range p.VulnsReports {
		merged.CriticalCount += report.Summary.CriticalCount
		merged.HighCount += report.Summary.HighCount
		merged.MediumCount += report.Summary.MediumCount
		merged.LowCount += report.Summary.LowCount
		merged.UnknownCount += report.Summary.UnknownCount
	}
	return merged
}

//line pkg/report/templates/workload_report.qtpl:22
func (p *WorkloadReport) StreamBody(qw422016 *qt422016.Writer) {
//line pkg/report/templates/workload_report.qtpl:22
	qw422016.N().S(`
  <style>
  a {
    color: inherit;
  }
  @media print {
    .container {
        display: inline;
    }
  }
  </style>
  <div class="container border-right border-left" style="height: 100%; overflow: scroll;">
    <div class="col mt-5">
      <div class="row text-center">
        `)
//line pkg/report/templates/workload_report.qtpl:36
	streamimgAquaLogo(qw422016)
//line pkg/report/templates/workload_report.qtpl:36
	qw422016.N().S(`
      </div>
      <div class="row mt-4 text-center">
        <h2 class="text-muted mx-auto">Aqua Starboard Workload Security Report</h2>
      </div>
      <div class="row text-center">
        <h3 class="text-muted mx-auto">Workload: `)
//line pkg/report/templates/workload_report.qtpl:42
	qw422016.E().V(p.Workload.Kind)
//line pkg/report/templates/workload_report.qtpl:42
	qw422016.N().S(`/`)
//line pkg/report/templates/workload_report.qtpl:42
	qw422016.E().S(p.Workload.Name)
//line pkg/report/templates/workload_report.qtpl:42
	qw422016.N().S(`</h3>
      </div>
      <div class="row text-center">
        <h3 class="text-muted mx-auto">Namespace: `)
//line pkg/report/templates/workload_report.qtpl:45
	qw422016.E().S(p.Workload.Namespace)
//line pkg/report/templates/workload_report.qtpl:45
	qw422016.N().S(`</h3>
      </div>
      <div class="row text-center">
        <h3 class="text-muted mx-auto">Generated on `)
//line pkg/report/templates/workload_report.qtpl:48
	qw422016.E().S(p.GeneratedAt.Format("2 Jan 2006 15:04:01"))
//line pkg/report/templates/workload_report.qtpl:48
	qw422016.N().S(`</h3>
      </div>

      <div class="row mt-5 px-3">
        <h4>Table Of Contents</h4>
      </div>

                <div class="row">
                    <ul>
                        `)
//line pkg/report/templates/workload_report.qtpl:57
	if len(p.VulnsReports) > 0 {
//line pkg/report/templates/workload_report.qtpl:57
		qw422016.N().S(`
                        <li>
                            <a href="#vuln_header">Vulnerabilities</a></li>
                            <ul>
                              `)
//line pkg/report/templates/workload_report.qtpl:61
		for container, _ := range p.VulnsReports {
//line pkg/report/templates/workload_report.qtpl:61
			qw422016.N().S(`
                                <li><a href="#vulns_container_`)
//line pkg/report/templates/workload_report.qtpl:62
			qw422016.E().S(container)
//line pkg/report/templates/workload_report.qtpl:62
			qw422016.N().S(`">`)
//line pkg/report/templates/workload_report.qtpl:62
			qw422016.E().S(container)
//line pkg/report/templates/workload_report.qtpl:62
			qw422016.N().S(`</a></li>
                              `)
//line pkg/report/templates/workload_report.qtpl:63
		}
//line pkg/report/templates/workload_report.qtpl:63
		qw422016.N().S(`
                            </ul>
                        </li>
                        `)
//line pkg/report/templates/workload_report.qtpl:66
	}
//line pkg/report/templates/workload_report.qtpl:66
	qw422016.N().S(`
                        `)
//line pkg/report/templates/workload_report.qtpl:67
	if p.ConfigAuditReport != nil && len(p.ConfigAuditReport.Report.PodChecks) > 0 {
//line pkg/report/templates/workload_report.qtpl:67
		qw422016.N().S(`
                        <li>
                            <a href="#ca_header">Configuration Audit</a>
                            <ul>
                              <li><a href="#ca_pod_checks">Pod Checks</a></li>
                                `)
//line pkg/report/templates/workload_report.qtpl:72
		for container, _ := range p.ConfigAuditReport.Report.ContainerChecks {
//line pkg/report/templates/workload_report.qtpl:72
			qw422016.N().S(`
                                  <li><a href="#ca_container_`)
//line pkg/report/templates/workload_report.qtpl:73
			qw422016.E().S(container)
//line pkg/report/templates/workload_report.qtpl:73
			qw422016.N().S(`">`)
//line pkg/report/templates/workload_report.qtpl:73
			qw422016.E().S(container)
//line pkg/report/templates/workload_report.qtpl:73
			qw422016.N().S(`</a></li>
                                `)
//line pkg/report/templates/workload_report.qtpl:74
		}
//line pkg/report/templates/workload_report.qtpl:74
		qw422016.N().S(`
                            </ul>
                        </li>
                        `)
//line pkg/report/templates/workload_report.qtpl:77
	}
//line pkg/report/templates/workload_report.qtpl:77
	qw422016.N().S(`
                    </ul>
                </div>


                `)
//line pkg/report/templates/workload_report.qtpl:82
	if len(p.VulnsReports) > 0 {
//line pkg/report/templates/workload_report.qtpl:82
		qw422016.N().S(`
                <!-- Vulnerabilities -->
                <div class="row text-center border-bottom mt-4">
                    <h3 class="mx-auto " id="vuln_header" style="color: rgb(0, 160, 170);">Vulnerabilities</h3>
                </div>
                <!-- Cards -->
                <div class="">
                    <div class="row my-5" style="font-size:small;">
                        <!-- Scanner -->
                        <div class="col-3 border rounded shadow px-3 py-2 ml-4 ">
                            <div class="row text-center">
                                <div class="col">
                                    <p class="mb-2 pb-1 border-bottom">Scanner</p>
                                </div>
                             </div>
                             <div class="row">
                                <div class="col">
                                `)
//line pkg/report/templates/workload_report.qtpl:100
		var scanner_name, scanner_vendor, scanner_version, creation_timestamp string
		for _, report := range p.VulnsReports {
			scanner_name = report.Scanner.Name
			scanner_vendor = report.Scanner.Vendor
			scanner_version = report.Scanner.Version
			creation_timestamp = report.UpdateTimestamp.Format("2 Jan 2006 15:04:01")
			break
		}

//line pkg/report/templates/workload_report.qtpl:108
		qw422016.N().S(`
                                    <p class="my-0">Name:  `)
//line pkg/report/templates/workload_report.qtpl:109
		qw422016.E().S(scanner_name)
//line pkg/report/templates/workload_report.qtpl:109
		qw422016.N().S(`</p>
                                    <p class="my-0">Vendor:  `)
//line pkg/report/templates/workload_report.qtpl:110
		qw422016.E().S(scanner_vendor)
//line pkg/report/templates/workload_report.qtpl:110
		qw422016.N().S(`</p>
                                    <p class="my-0">Version:  `)
//line pkg/report/templates/workload_report.qtpl:111
		qw422016.E().S(scanner_version)
//line pkg/report/templates/workload_report.qtpl:111
		qw422016.N().S(`</p>
                                </div>
                             </div>
                        </div>
                        <!-- summary -->
                        <div class="col-5 border rounded shadow py-2 mx-auto ">
                            <div class="row text-center">
                               <div class="col">
                                   <p class="mb-2 pb-1 border-bottom">Summary</p>
                               </div>
                            </div>
                            <div class="row">
                                `)
//line pkg/report/templates/workload_report.qtpl:124
		summary := p.GetMergedVulnsSummary()

//line pkg/report/templates/workload_report.qtpl:125
		qw422016.N().S(`
                                `)
//line pkg/report/templates/workload_report.qtpl:126
		if summary.CriticalCount > 0 {
//line pkg/report/templates/workload_report.qtpl:126
			qw422016.N().S(`
                                <div class="col text-center p-0 text-danger font-weight-bold">
                                `)
//line pkg/report/templates/workload_report.qtpl:128
		} else {
//line pkg/report/templates/workload_report.qtpl:128
			qw422016.N().S(`
                                <div class="col text-center p-0">
                                `)
//line pkg/report/templates/workload_report.qtpl:130
		}
//line pkg/report/templates/workload_report.qtpl:130
		qw422016.N().S(`
                                    <p class="mx-auto mb-1">`)
//line pkg/report/templates/workload_report.qtpl:131
		qw422016.N().D(summary.CriticalCount)
//line pkg/report/templates/workload_report.qtpl:131
		qw422016.N().S(`</p>
                                    <p class="mx-auto ">CRITICAL</p>
                                </div>
                                `)
//line pkg/report/templates/workload_report.qtpl:134
		if summary.HighCount > 0 {
//line pkg/report/templates/workload_report.qtpl:134
			qw422016.N().S(`
                                <div class="col text-center p-0 text-danger font-weight-bold">
                                `)
//line pkg/report/templates/workload_report.qtpl:136
		} else {
//line pkg/report/templates/workload_report.qtpl:136
			qw422016.N().S(`
                                <div class="col text-center p-0">
                                `)
//line pkg/report/templates/workload_report.qtpl:138
		}
//line pkg/report/templates/workload_report.qtpl:138
		qw422016.N().S(`
                                    <p class="mx-auto mb-1">`)
//line pkg/report/templates/workload_report.qtpl:139
		qw422016.N().D(summary.HighCount)
//line pkg/report/templates/workload_report.qtpl:139
		qw422016.N().S(`</p>
                                    <p class="mx-auto ">HIGH</p>
                                </div>
                                `)
//line pkg/report/templates/workload_report.qtpl:142
		if summary.MediumCount > 0 {
//line pkg/report/templates/workload_report.qtpl:142
			qw422016.N().S(`
                                <div class="col text-center p-0 text-warning font-weight-bold">
                                `)
//line pkg/report/templates/workload_report.qtpl:144
		} else {
//line pkg/report/templates/workload_report.qtpl:144
			qw422016.N().S(`
                                <div class="col text-center p-0">
                                `)
//line pkg/report/templates/workload_report.qtpl:146
		}
//line pkg/report/templates/workload_report.qtpl:146
		qw422016.N().S(`
                                    <p class="mx-auto mb-1">`)
//line pkg/report/templates/workload_report.qtpl:147
		qw422016.N().D(summary.MediumCount)
//line pkg/report/templates/workload_report.qtpl:147
		qw422016.N().S(`</p>
                                    <p class="mx-auto ">MEDIUM</p>
                                </div>
                                <div class="col text-center p-0">
                                    <p class="mx-auto mb-1">`)
//line pkg/report/templates/workload_report.qtpl:151
		qw422016.N().D(summary.LowCount)
//line pkg/report/templates/workload_report.qtpl:151
		qw422016.N().S(`</p>
                                    <p class="mx-auto ">LOW</p>
                                </div>
                                <div class="col text-center p-0">
                                    <p class="mx-auto mb-1">`)
//line pkg/report/templates/workload_report.qtpl:155
		qw422016.N().D(summary.UnknownCount)
//line pkg/report/templates/workload_report.qtpl:155
		qw422016.N().S(`</p>
                                    <p class="mx-auto ">UNKNOWN</p>
                                </div>
                            </div>
                        </div>
                        <!-- Metadata -->
                        <div class="col-3 border rounded shadow px-3 py-2 mr-4">
                            <div class="row text-center">
                                <div class="col">
                                    <p class="mb-2 pb-1 border-bottom">Metadata</p>
                                </div>
                             </div>
                             <div class="row">
                                <div class="col">
                                    <p class="my-0">
                                        Generated at:  `)
//line pkg/report/templates/workload_report.qtpl:170
		qw422016.E().S(creation_timestamp)
//line pkg/report/templates/workload_report.qtpl:170
		qw422016.N().S(`
                                    </p>
                                </div>
                             </div>
                        </div>

                    </div>      
                </div>
                `)
//line pkg/report/templates/workload_report.qtpl:178
	}
//line pkg/report/templates/workload_report.qtpl:178
	qw422016.N().S(`
                
                `)
//line pkg/report/templates/workload_report.qtpl:180
	for container, report := range p.VulnsReports {
//line pkg/report/templates/workload_report.qtpl:180
		qw422016.N().S(`
                
                  <div class="row"><h5 class="text-info" id="vulns_container_`)
//line pkg/report/templates/workload_report.qtpl:182
		qw422016.E().S(container)
//line pkg/report/templates/workload_report.qtpl:182
		qw422016.N().S(`">Container `)
//line pkg/report/templates/workload_report.qtpl:182
		qw422016.E().S(container)
//line pkg/report/templates/workload_report.qtpl:182
		qw422016.N().S(`</h5></div>
                  <div class="row"><p>`)
//line pkg/report/templates/workload_report.qtpl:183
		qw422016.E().S(report.Registry.Server)
//line pkg/report/templates/workload_report.qtpl:183
		qw422016.N().S(`/`)
//line pkg/report/templates/workload_report.qtpl:183
		qw422016.E().S(report.Artifact.Repository)
//line pkg/report/templates/workload_report.qtpl:183
		qw422016.N().S(`:`)
//line pkg/report/templates/workload_report.qtpl:183
		qw422016.E().S(report.Artifact.Tag)
//line pkg/report/templates/workload_report.qtpl:183
		qw422016.N().S(`</p></div>
                  `)
//line pkg/report/templates/workload_report.qtpl:184
		if len(report.Vulnerabilities) == 0 {
//line pkg/report/templates/workload_report.qtpl:184
			qw422016.N().S(`
                    <div class="row">
                      <p class="alert alert-success py-0 m-0" style="font-size: small;">No Vulnerabilities</p>
                    </div>                  
                  `)
//line pkg/report/templates/workload_report.qtpl:188
		} else {
//line pkg/report/templates/workload_report.qtpl:188
			qw422016.N().S(`

                  <div class="row">
                    <table class="table table-sm table-bordered">
                      <thead>
                        <tr>
                          <th scope="col">ID</th>
                          <th scope="col">Severity</th>
                          <th scope="col">Resource</th>
                          <th scope="col">Installed Version</th>
                          <th scope="col">Fixed Version</th>
                        </tr>
                      </thead>
                      <tbody>
                        `)
//line pkg/report/templates/workload_report.qtpl:202
			for _, v := range report.Vulnerabilities {
//line pkg/report/templates/workload_report.qtpl:202
				qw422016.N().S(`
                        <tr>
                          <td>
                            <a target="_blank" href="`)
//line pkg/report/templates/workload_report.qtpl:205
				qw422016.E().S(v.PrimaryLink)
//line pkg/report/templates/workload_report.qtpl:205
				qw422016.N().S(`">`)
//line pkg/report/templates/workload_report.qtpl:205
				qw422016.E().S(v.VulnerabilityID)
//line pkg/report/templates/workload_report.qtpl:205
				qw422016.N().S(`</a>
                          </td>
                          <td>`)
//line pkg/report/templates/workload_report.qtpl:207
				qw422016.E().V(v.Severity)
//line pkg/report/templates/workload_report.qtpl:207
				qw422016.N().S(`</td>
                          <td>`)
//line pkg/report/templates/workload_report.qtpl:208
				qw422016.E().S(v.Resource)
//line pkg/report/templates/workload_report.qtpl:208
				qw422016.N().S(`</td>
                          <td>`)
//line pkg/report/templates/workload_report.qtpl:209
				qw422016.E().S(v.InstalledVersion)
//line pkg/report/templates/workload_report.qtpl:209
				qw422016.N().S(`</td>
                          <td>`)
//line pkg/report/templates/workload_report.qtpl:210
				qw422016.E().S(v.FixedVersion)
//line pkg/report/templates/workload_report.qtpl:210
				qw422016.N().S(`</td>
                        </tr>
                        `)
//line pkg/report/templates/workload_report.qtpl:212
			}
//line pkg/report/templates/workload_report.qtpl:212
			qw422016.N().S(`
                      </tbody>
                    </table>
                  </div>
                `)
//line pkg/report/templates/workload_report.qtpl:216
		}
//line pkg/report/templates/workload_report.qtpl:216
		qw422016.N().S(`
                `)
//line pkg/report/templates/workload_report.qtpl:217
	}
//line pkg/report/templates/workload_report.qtpl:217
	qw422016.N().S(`

                <!-- Config Audits -->
                `)
//line pkg/report/templates/workload_report.qtpl:220
	if p.ConfigAuditReport != nil && len(p.ConfigAuditReport.Report.PodChecks) > 0 {
//line pkg/report/templates/workload_report.qtpl:220
		qw422016.N().S(`
                  <div class="row pt-3 text-center border-bottom my-4">
                      <h3 class="mx-auto" id="ca_header" style="color: rgb(0, 160, 170);">Configuration Audit</h3>
                  </div>
                  <!-- Cards -->
                <div class="">
                    <div class="row my-5" style="font-size:small;">
                        <!-- Scanner -->
                        <div class="col-3 border rounded shadow px-3 py-2 ml-4 ">
                            <div class="row text-center">
                                <div class="col">
                                    <p class="mb-2 pb-1 border-bottom">Scanner</p>
                                </div>
                             </div>
                             <div class="row">
                                <div class="col">
                                    <p class="my-0">Name:  `)
//line pkg/report/templates/workload_report.qtpl:236
		qw422016.E().S(p.ConfigAuditReport.Report.Scanner.Name)
//line pkg/report/templates/workload_report.qtpl:236
		qw422016.N().S(`</p>
                                    <p class="my-0">Vendor:  `)
//line pkg/report/templates/workload_report.qtpl:237
		qw422016.E().S(p.ConfigAuditReport.Report.Scanner.Vendor)
//line pkg/report/templates/workload_report.qtpl:237
		qw422016.N().S(`</p>
                                    <p class="my-0">Version:  `)
//line pkg/report/templates/workload_report.qtpl:238
		qw422016.E().S(p.ConfigAuditReport.Report.Scanner.Version)
//line pkg/report/templates/workload_report.qtpl:238
		qw422016.N().S(`</p>
                                </div>
                             </div>
                        </div>
                        <!-- Summary -->
                        <div class="col-3 border rounded shadow py-2 mx-auto ">
                            <div class="row text-center">
                               <div class="col">
                                   <p class="mb-2 pb-1 border-bottom">Summary</p>
                               </div>
                            </div>
                            <div class="row">
                              `)
//line pkg/report/templates/workload_report.qtpl:251
		sumCritical := p.ConfigAuditReport.Report.Summary.CriticalCount
		sumHigh := p.ConfigAuditReport.Report.Summary.HighCount
		sumMedium := p.ConfigAuditReport.Report.Summary.MediumCount
		sumLow := p.ConfigAuditReport.Report.Summary.LowCount

//line pkg/report/templates/workload_report.qtpl:255
		qw422016.N().S(`

                              <div class="col text-center p-0 text-danger font-weight-bold">
                                <p class="mx-auto mb-1">`)
//line pkg/report/templates/workload_report.qtpl:258
		qw422016.N().D(sumCritical)
//line pkg/report/templates/workload_report.qtpl:258
		qw422016.N().S(`</p>
                                <p class="mx-auto">CRITICAL</p>
                              </div>

                              <div class="col text-center p-0 text-danger font-weight-bold">
                                <p class="mx-auto mb-1">`)
//line pkg/report/templates/workload_report.qtpl:263
		qw422016.N().D(sumHigh)
//line pkg/report/templates/workload_report.qtpl:263
		qw422016.N().S(`</p>
                                <p class="mx-auto">HIGH</p>
                              </div>

                              <div class="col text-center p-0 text-warning font-weight-bold">
                                <p class="mx-auto mb-1">`)
//line pkg/report/templates/workload_report.qtpl:268
		qw422016.N().D(sumMedium)
//line pkg/report/templates/workload_report.qtpl:268
		qw422016.N().S(`</p>
                                <p class="mx-auto">MEDIUM</p>
                              </div>

                              <div class="col text-center p-0">
                                <p class="mx-auto mb-1">`)
//line pkg/report/templates/workload_report.qtpl:273
		qw422016.N().D(sumLow)
//line pkg/report/templates/workload_report.qtpl:273
		qw422016.N().S(`</p>
                                <p class="mx-auto">LOW</p>
                              </div>
                            </div>
                        </div>
                        <!-- Metadata -->
                        <div class="col-3 border rounded shadow px-3 py-2 mr-4">
                            <div class="row text-center">
                                <div class="col">
                                    <p class="mb-2 pb-1 border-bottom">Metadata</p>
                                </div>
                             </div>
                             <div class="row">
                                <div class="col">
                                    <p class="my-0">
                                        Generated at:  `)
//line pkg/report/templates/workload_report.qtpl:288
		qw422016.E().S(p.ConfigAuditReport.Report.UpdateTimestamp.Format("2 Jan 2006 15:04:01"))
//line pkg/report/templates/workload_report.qtpl:288
		qw422016.N().S(`
                                    </p>
                                </div>
                             </div>
                        </div>
                    </div>
                </div>
                  <div class="row"><h5 class="text-info" id="ca_pod_checks">Pod Checks</h5></div>
                  <div class="row">
                      <table class="table table-sm table-bordered">
                          <thead>
                              <tr>
                                <th scope="col">PASS</th>
                                <th scope="col">ID</th>
                                <th scope="col">Severity</th>
                                <th scope="col">Category</th>
                              </tr>
                            </thead>
                            <tbody>
                              `)
//line pkg/report/templates/workload_report.qtpl:307
		for _, check := range p.ConfigAuditReport.Report.PodChecks {
//line pkg/report/templates/workload_report.qtpl:307
			qw422016.N().S(`
                                <tr>
                                  <td>`)
//line pkg/report/templates/workload_report.qtpl:309
			qw422016.E().V(check.Success)
//line pkg/report/templates/workload_report.qtpl:309
			qw422016.N().S(`</td>
                                  <td>`)
//line pkg/report/templates/workload_report.qtpl:310
			qw422016.E().S(check.ID)
//line pkg/report/templates/workload_report.qtpl:310
			qw422016.N().S(`</td>
                                  <td>`)
//line pkg/report/templates/workload_report.qtpl:311
			qw422016.E().V(check.Severity)
//line pkg/report/templates/workload_report.qtpl:311
			qw422016.N().S(`</td>
                                  <td>`)
//line pkg/report/templates/workload_report.qtpl:312
			qw422016.E().S(check.Category)
//line pkg/report/templates/workload_report.qtpl:312
			qw422016.N().S(`</td>
                                </tr>
                              `)
//line pkg/report/templates/workload_report.qtpl:314
		}
//line pkg/report/templates/workload_report.qtpl:314
		qw422016.N().S(`
                            </tbody>
                      </table>
                  </div>
                  `)
//line pkg/report/templates/workload_report.qtpl:318
		for container, checks := range p.ConfigAuditReport.Report.ContainerChecks {
//line pkg/report/templates/workload_report.qtpl:318
			qw422016.N().S(`
                    <div class="row"><h5 class="text-info" id="ca_container_`)
//line pkg/report/templates/workload_report.qtpl:319
			qw422016.E().S(container)
//line pkg/report/templates/workload_report.qtpl:319
			qw422016.N().S(`">Container `)
//line pkg/report/templates/workload_report.qtpl:319
			qw422016.E().S(container)
//line pkg/report/templates/workload_report.qtpl:319
			qw422016.N().S(`</h5></div>
                    <div class="row">
                        <table class="table table-sm table-bordered">
                            <thead>
                                <tr>
                                  <th scope="col">PASS</th>
                                  <th scope="col">ID</th>
                                  <th scope="col">Severity</th>
                                  <th scope="col">Category</th>
                                </tr>
                              </thead>
                              <tbody>
                                `)
//line pkg/report/templates/workload_report.qtpl:331
			for _, check := range checks {
//line pkg/report/templates/workload_report.qtpl:331
				qw422016.N().S(`
                                  <tr>
                                    <td>`)
//line pkg/report/templates/workload_report.qtpl:333
				qw422016.E().V(check.Success)
//line pkg/report/templates/workload_report.qtpl:333
				qw422016.N().S(`</td>
                                    <td>`)
//line pkg/report/templates/workload_report.qtpl:334
				qw422016.E().S(check.ID)
//line pkg/report/templates/workload_report.qtpl:334
				qw422016.N().S(`</td>
                                    <td>`)
//line pkg/report/templates/workload_report.qtpl:335
				qw422016.E().V(check.Severity)
//line pkg/report/templates/workload_report.qtpl:335
				qw422016.N().S(`</td>
                                    <td>`)
//line pkg/report/templates/workload_report.qtpl:336
				qw422016.E().S(check.Category)
//line pkg/report/templates/workload_report.qtpl:336
				qw422016.N().S(`</td>
                                  </tr>
                                `)
//line pkg/report/templates/workload_report.qtpl:338
			}
//line pkg/report/templates/workload_report.qtpl:338
			qw422016.N().S(`
                              </tbody>
                        </table>
                    </div>
                  `)
//line pkg/report/templates/workload_report.qtpl:342
		}
//line pkg/report/templates/workload_report.qtpl:342
		qw422016.N().S(`
                  `)
//line pkg/report/templates/workload_report.qtpl:343
	}
//line pkg/report/templates/workload_report.qtpl:343
	qw422016.N().S(`
            </div>
        </div>
`)
//line pkg/report/templates/workload_report.qtpl:346
}

//line pkg/report/templates/workload_report.qtpl:346
func (p *WorkloadReport) WriteBody(qq422016 qtio422016.Writer) {
//line pkg/report/templates/workload_report.qtpl:346
	qw422016 := qt422016.AcquireWriter(qq422016)
//line pkg/report/templates/workload_report.qtpl:346
	p.StreamBody(qw422016)
//line pkg/report/templates/workload_report.qtpl:346
	qt422016.ReleaseWriter(qw422016)
//line pkg/report/templates/workload_report.qtpl:346
}

//line pkg/report/templates/workload_report.qtpl:346
func (p *WorkloadReport) Body() string {
//line pkg/report/templates/workload_report.qtpl:346
	qb422016 := qt422016.AcquireByteBuffer()
//line pkg/report/templates/workload_report.qtpl:346
	p.WriteBody(qb422016)
//line pkg/report/templates/workload_report.qtpl:346
	qs422016 := string(qb422016.B)
//line pkg/report/templates/workload_report.qtpl:346
	qt422016.ReleaseByteBuffer(qb422016)
//line pkg/report/templates/workload_report.qtpl:346
	return qs422016
//line pkg/report/templates/workload_report.qtpl:346
}
