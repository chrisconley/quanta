# Analysis: DapperLib/Dapper

- **Repo**: DapperLib/Dapper (~18.2k stars)
- **Type**: .NET micro-ORM providing `IDbConnection` extension methods that map SQL results to typed objects (originally extracted from Stack Overflow).
- **Length**: ~21 KB. Long; dominated by a full BenchmarkDotNet results table and a stack of small usage vignettes.

## Sections

1. **H1 setext title + AppVeyor badge** — title rendered with `=` underline ("Dapper - a simple object mapper for .Net"), single CI badge beneath.
2. **Release Notes** — one-line pointer to the GitHub Releases page.
3. **Packages** — MyGet pre-release feed URL, then a six-row matrix table with columns NuGet Stable / NuGet Pre-release / Downloads / MyGet for Dapper, Dapper.EntityFramework, Dapper.EntityFramework.StrongName, Dapper.Rainbow, Dapper.SqlBuilder, Dapper.StrongName; followed by a bulleted "Package Purposes" list that re-describes each package in plain prose.
4. **Sponsors** — origin statement ("originally developed for and by Stack Overflow"), shout-out to Dapper Plus (paid add-on) and AWS (.NET on AWS Open Source Software Fund), plus a 728x90 Dapper Plus banner `<img>`.
5. **Features** — one-paragraph pitch followed by the three core APIs in a single `csharp` block (`Execute`, `Query<T>`, `QuerySingle<T>`) and a three-bullet list of accepted `args` shapes (POCO, `Dictionary<string,object>`, `DynamicParameters`).
6. **Execute a query and map it to a list of typed objects** — `Dog` class + `Query<Dog>` + xUnit `Assert.Equal` lines.
7. **Execute a query and map it to a list of dynamic objects** — `.AsList()` on a `Query(...)` result.
8. **Execute a Command that returns no results** — includes `set nocount on`/`drop table #t` noise to show realistic batch SQL.
9. **Execute a Command multiple times** — bulk insert via array-of-anonymous-types; second vignette showing it with an existing `List<Foo>`.
10. **Performance** — motivation paragraph, `dotnet run --project ... --join` command, fenced `ini` block with hardware spec, then a ~50-row BenchmarkDotNet markdown table comparing Dapper to Hand Coded, SqlMarshal, Mighty, LINQ to DB, RepoDB, Norm, ServiceStack, Massive, DevExpress.XPO, Belgrade, EF Core, EF 6, NHibernate. Footer invites ORM-contribution patches and points at RawDataAccessBencher / OrmBenchmark.
11. **Parameterized queries** — anonymous-class parameters, then `DynamicParameters` with `queryParams.Add("param1", value, DbType.Guid)` and `AddDynamicParams`.
12. **List Support** — `IEnumerable<int>` auto-expansion; shows the before-and-after SQL rewrite (`in @Ids` -> `in (@Ids1, @Ids2, @Ids3)`).
13. **Literal replacements** — `{=Admin}` syntax and a caveat about query plans / `OPTIMIZE FOR UNKNOWN`.
14. **Buffered vs Unbuffered readers** — default buffering rationale and `buffered: false`.
15. **Multi Mapping** — full walk-through with `Post`/`User` classes, the `(post, user) => { post.Owner = user; return post; }` mapper, the `<Post, User, Post>` type-arg explanation, and a note on `splitOn`.
16. **Multiple Results** — `QueryMultiple` + `multi.Read<T>()` usage.
17. **Stored Procedures** — `CommandType.StoredProcedure` basic and advanced (`ParameterDirection.Output`/`ReturnValue`).
18. **Ansi Strings and varchar** — `DbString` with `IsFixedLength`, `Length`, `IsAnsi`.
19. **Type Switching Per Row** — `IDataReader.GetRowParser` with a `Shapes` table and `Circle`/`Square`/`Triangle` switch.
20. **User Defined Variables in MySQL/MariaDB** — connection-string flag `Allow User Variables=True`.
21. **Limitations and caveats** — `ConcurrentDictionary` query cache, "worries about the 95% scenario."
22. **Will Dapper work with my DB provider?** — lists SQLite, SQL CE, Firebird, Oracle, MariaDB, MySQL, PostgreSQL, SQL Server.
23. **Do you have a comprehensive list of examples?** — pointer to `tests/Dapper.Tests`.
24. **Who is using this?** — one line: Stack Overflow.

## Tone & style

- **Register**: Reference-manual-meets-cookbook. Terse, declarative, example-first; each feature gets a 1-3 sentence setup and a code block.
- **Voice**: Third-person describing the library ("Dapper caches information...", "Dapper allows you to...") with occasional second-person imperatives for usage.
- **Formatting**: setext-underline headings (`===` for H1, `---` for H2) throughout instead of ATX `#`; `csharp` fenced blocks dominate; exactly three markdown tables (packages matrix, benchmark results, none inline); xUnit `Assert.Equal` lines are retained in examples instead of `Console.WriteLine`-style output.

## Notable conventions

- **Benchmark table is the centerpiece and Dapper is not #1** — Hand Coded `SqlCommand` beats Dapper, and the table is left unsorted/unfiltered. Honest rather than cherry-picked; "A key feature of Dapper is performance" is followed by data that shows Dapper is fast-among-ORMs, not fastest-overall.
- **Benchmarks ship with a command to reproduce them** — `dotnet run --project ...\Dapper.Tests.Performance\ -c Release -f net8.0 -- -f * --join`, plus `<kbd>Ctrl</kbd>+<kbd>F5</kbd>` guidance and pointers to two competing benchmark suites.
- **xUnit assertions as documentation** — examples end with `Assert.Equal(...)` rather than comments or print statements, implying the examples double as executable tests.
- **Setext headings everywhere** — uses underline-style H1/H2 (`---` below the heading) rather than `##`, a carryover from the project's 2011 README era that has never been modernized.
- **"Package Purposes" duplicates the table** — the matrix has package names as links and the bullet list re-lists them with purpose; a hint the table was bolted on later and the prose wasn't deleted.
- **Sponsor section with a commercial-product banner** — Dapper Plus, a paid product on a separate domain, gets a full-width banner image inline between "Sponsors" and "Features"; reads as paid placement but disclosed.
- **"Who is using this?" answered with one sentence** — "Dapper is in production use at Stack Overflow" is the entire section, and that's enough given the origin story.
- **Typo preserved** — "anonyomous types" in the Features section has not been fixed, consistent with the overall "README has accreted over a decade" feel.

## Takeaway pattern

This is the "old, practical, still-shipping" README archetype: a feature-by-feature cookbook that grew one vignette at a time over ~15 years, anchored by a big honest benchmark table and a one-sentence social proof ("in production at Stack Overflow"). It works because Dapper's surface is small enough that a linear list of 20 recipes covers it, and because the tone never oversells — the benchmark puts Hand-Coded SQL ahead of Dapper without comment, and the "Limitations" section admits the library deliberately ignores 5% of ORM use cases. Lesson: when your value prop is "small, fast, boring," a README that looks small, fast, and a little boring is on-brand; a marketing-forward rewrite would undercut the positioning.
